package trackdataserver

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	config "track-server-api/config"
	pb "track-server-api/rpc"
)

type StorageConnection struct {
	AccountId          string `json:"accountId"`
	ApiUrl             string `json:"apiUrl"`
	AuthorizationToken string `json:"authorizationToken"`
	DownloadUrl        string `json:"downloadUrl"`
	MinimumPartSize    int    `json:"minimumPartSize"`
}

type UploadUrl struct {
	BucketId           string `json:"bucketID"`
	UploadUrl          string `json:uploadUrl`
	AuthorizationToken string `json:authorizationToken`
}

type StorageFileInfo struct {
	FileId string `json:"fileId"`
}

func OpenStorageConnection() (*StorageConnection, error) {

	var client = &http.Client{Timeout: config.StorageConfig.Timeout * time.Second}
	req, err := http.NewRequest("GET", config.StorageConfig.AuthEndpoint, nil)
	req.SetBasicAuth(config.StorageConfig.AccountId, config.StorageConfig.Key)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("HTTP error %d while opening storage connection", resp.StatusCode)
		err := errors.New(msg)
		return nil, err
	}

	var sc = new(StorageConnection)
	err = json.Unmarshal(body, &sc)
	if err != nil {
		return nil, err
	}

	return sc, nil
}

func GetTrackChunkFromStorage(storageId string, trackChunkPB *pb.TrackChunk, sc *StorageConnection) (*pb.TrackChunk, error) {

	endpoint := fmt.Sprintf("%s%s%s", sc.DownloadUrl, config.StorageConfig.FileEndpoint, storageId)
	rangeMsg := fmt.Sprintf("bytes=%d-%d", trackChunkPB.StartPosition, trackChunkPB.StartPosition+trackChunkPB.NumBytes-1)

	var client = &http.Client{Timeout: config.StorageConfig.Timeout * time.Second}
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Authorization", sc.AuthorizationToken)
	req.Header.Set("Range", rangeMsg)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusPartialContent {
		if resp.StatusCode == http.StatusRequestedRangeNotSatisfiable {
			return nil, io.EOF
		}
		msg := fmt.Sprintf("HTTP error %d while getting track data from storage", resp.StatusCode)
		err := errors.New(msg)
		return nil, err
	}

	td := &pb.TrackChunk{
		StartPosition: trackChunkPB.StartPosition,
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	written, err := io.Copy(writer, resp.Body)
	if err != nil {
		return nil, err
	}

	td.Data = b.Bytes()
	td.NumBytes = int32(written)
	return td, nil
}

func GetUploadUrl(sc *StorageConnection) (*UploadUrl, error) {

	var client = &http.Client{Timeout: config.StorageConfig.Timeout * time.Second}

	endpoint := fmt.Sprintf("%s%s", sc.ApiUrl, config.StorageConfig.UploadEndpoint)
	s := fmt.Sprintf(`{"bucketId": "%s"}`, config.StorageConfig.BucketId)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(s)))
	req.Header.Set("Authorization", sc.AuthorizationToken)

	fmt.Printf("GetUploadUrl: %v\n%v\n", req.URL, req.Header)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("HTTP error %d response to getting upload URL", resp.StatusCode)
		err := errors.New(msg)
		return nil, err
	}

	var uploadURL = new(UploadUrl)
	err = json.Unmarshal(body, &uploadURL)
	if err != nil {
		return nil, err
	}

	return uploadURL, nil
}

func UploadTrackToStorage(trackUpload *pb.TrackUpload, uploadUrl *UploadUrl, sc *StorageConnection) (*StorageFileInfo, error) {

	var client = &http.Client{Timeout: config.StorageConfig.Timeout * time.Second} // with large upload this timeout has to be big enough...

	h := sha1.New()
	h.Write(trackUpload.Data)
	sha := fmt.Sprintf("%x", h.Sum(nil))

	req, err := http.NewRequest("POST", uploadUrl.UploadUrl, bytes.NewBuffer(trackUpload.Data))
	req.Header.Set("Authorization", uploadUrl.AuthorizationToken)
	req.Header.Set("X-Bz-File-Name", trackUpload.Name)
	req.Header.Set("Content-Type", "audio/aac")
	req.Header.Set("X-Bz-Content-Sha1", sha)
	req.Header.Set("X-Bz-Info-Author", "unknown")

	fmt.Printf("UploadTrack: %v\n%v\n", req.URL, req.Header)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("HTTP error %d response to upload POST", resp.StatusCode)
		err := errors.New(msg)
		return nil, err
	}

	var fileInfo = new(StorageFileInfo)
	err = json.Unmarshal(body, &fileInfo)
	if err != nil {
		return nil, err
	}

	return fileInfo, nil
}
