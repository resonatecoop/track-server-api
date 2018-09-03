package trackdataserver

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	pb "track-server-api/rpc"
)

type StorageConnection struct {
	AccountId          string `json:"accountId"`
	ApiUrl             string `json:"apiUrl"`
	AuthorizationToken string `json:"authorizationToken"`
	DownloadUrl        string `json:"downloadUrl"`
	MinimumPartSize    string `json:"minimumPartSize"`
}

func OpenStorageConnection() (*StorageConnection, error) {
	var client = &http.Client{Timeout: 5 * time.Second}

	url := fmt.Sprintf("https://api.backblazeb2.com/b2api/v1/b2_authorize_account -u \"%d:%d\"", 1, 2)

	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("HTTP error %d while opening storage connection", resp.StatusCode)
		err := errors.New(msg)
		return nil, err
	}

	sc := &StorageConnection{}
	err = json.NewDecoder(resp.Body).Decode(sc)
	if err != nil {
		return nil, err
	}
	return sc, nil
}

func GetTrackChunkFromStorage(trackServerId string, trackChunkPB *pb.TrackChunk, sc *StorageConnection) (*pb.TrackChunk, error) {
	endpoint := fmt.Sprintf("/b2api/v1/b2_download_file_by_id?fileId=\"%s\"", trackServerId)
	url := fmt.Sprintf("%s/%s", endpoint, trackServerId)
	rangeMsg := fmt.Sprintf("%d-%d", 0, trackChunkPB.NumBytes)

	var client = &http.Client{Timeout: 5 * time.Second}
	println(url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("AuthorizationToken", sc.AuthorizationToken)
	req.Header.Set("Range", rangeMsg)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
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
