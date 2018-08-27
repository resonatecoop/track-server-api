package playserver

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"time"
	"errors"
	"bufio"

	pb "track-server-api/rpc"
)

type StorageConnection struct {
	AccountId string `json:"accountId"`
	ApiUrl string `json:"apiUrl"`
	AuthorizationToken string `json:"authorizationToken"`
	DownloadUrl string `json:"downloadUrl"`
	MinimumPartSize string `json:"minimumPartSize"`
}

func openConnection() (*StorageConnection, error) {
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

func getTrackData(trackDataPB *pb.TrackData, sc *StorageConnection) (*pb.TrackData, error) {
	endpoint := fmt.Sprintf("/b2api/v1/b2_download_file_by_id?fileId=\"%s\"",trackDataPB.TrackServerId)
	url := fmt.Sprintf("%s/%s", endpoint, trackDataPB.TrackServerId)

	rangeMsg := fmt.Sprintf("%d-%d", 0, trackDataPB.NumBytes)

	var client = &http.Client{Timeout: 5 * time.Second}

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

	td := &pb.TrackData{
		TrackServerId: trackDataPB.TrackServerId,
		StartPosition: trackDataPB.StartPosition,
	}
	
	writer := bufio.NewWriter(td.Data)

	written, err := io.Copy(writer, resp.Body)
	if err != nil {
		return nil, err
	}

	td.NumBytes = int32(written)
	return td, nil
}
