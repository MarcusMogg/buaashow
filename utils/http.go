package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// Post 发起一个post请求
func Post(url string, contentType string, data string) (res []byte, err error) {
	client := http.Client{
		Timeout: time.Second,
	}

	resp, err := client.Post(url, contentType, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	res, err = ioutil.ReadAll(resp.Body)
	return
}
