package trackdataserver

import (
	"bufio"
	"bytes"
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

func OpenStorageConnection() (*StorageConnection, error) {

	var client = &http.Client{Timeout: 5 * time.Second}
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

	var client = &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Authorization", sc.AuthorizationToken)
	req.Header.Set("Range", rangeMsg)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusPartialContent {
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
