package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 5 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	p := make([]byte, 10)

	for {
		n, err := r.Body.Read(p)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		fmt.Println(string(p[:n]))
	}
	return json.NewDecoder(r.Body).Decode(target)
}

type Foo struct {
	Bar    string `json:"message"`
	Status int    `json:"status"`
}

type people struct {
	Number int `json:"number"`
}

func main() {
	foo1 := new(Foo) // or &Foo{}

	s := fmt.Sprintf("https://api.backblazeb2.com/b2api/v1/b2_authorize_account -u \"%d:%d\"", 1, 2)
	getJson(s, &foo1)
	println(foo1.Bar)
	println(foo1.Status)
}
