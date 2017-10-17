package main

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"runtime"
)
import _ "net/http/pprof"

type ReqToken struct {
	App_key    string `json:"app_key"`
	App_secret string `json:"app_secret"`
}

type RespToken struct {
	Access_token string `json:"access_token"`
}

func CallbackTokenRedirect(resp http.ResponseWriter, req *http.Request) {
	req_body, err := ioutil.ReadAll(req.Body)
	log.Println(string(req_body))
	if err != nil {
		return
	}
	var resp_token ReqToken
	err = json.Unmarshal(req_body, &resp_token)
	if err != nil {
		return
	}

	resp.Write([]byte("success"))
}

func CallbackTokenGet(resp http.ResponseWriter, req *http.Request) {
	req_body, err := ioutil.ReadAll(req.Body)
	log.Println(string(req_body))
	if err != nil {
		return
	}

	var req_token ReqToken
	err = json.Unmarshal(req_body, &req_token)
	if err != nil {
		return
	}

	resp.Write(req_body)
}

func main() {
	http.HandleFunc("/token_redirect", CallbackTokenRedirect)
	http.HandleFunc("/token_get", CallbackTokenGet)
	_ = http.ListenAndServe("127.0.0.1:8101", nil)
	return
}
