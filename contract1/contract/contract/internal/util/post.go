package util

import (
	"bytes"
	"contract/internal/config"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Info struct {
	Timestamp int    `json:"timestamp"`
	Status    int    `json:"status"`
	Error     error  `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

// Post 发起Post请求
func Post(bytesData []byte) error {
	resp, err := http.Post(config.Setting.Java.MintUrl, config.Setting.Java.Type, bytes.NewReader(bytesData))
	if err != nil {
		Logger.Error(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Logger.Error(err.Error())
	}
	var info Info
	json.Unmarshal(body, &info)
	log.Println(info.Status)
	if info.Status != 0 {
		return info.Error
	}
	return nil
}
