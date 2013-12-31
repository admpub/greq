// greq - simple http request library
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package greq

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var Host = "http://0.0"

// Get the path from the Server.
func Get(path string) ([]byte, *http.Response) {
	return Request("GET", Host+path, nil)
}

// Post data to the path on the Server.
func Post(path string, data interface{}) ([]byte, *http.Response) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, nil
	}
	js := bytes.NewBuffer(b)
	return Request("POST", Host+path, js)
}

// Put data to the path on the Server.
func Put(path string, data interface{}) ([]byte, *http.Response) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, nil
	}
	js := bytes.NewBuffer(b)
	return Request("PUT", Host+path, js)
}

// Send delete to the path on the Server.
func Delete(path string) ([]byte, *http.Response) {
	return Request("DELETE", Host+path, nil)
}

// Generic Request method
func Request(method string, url string, body io.Reader) ([]byte, *http.Response) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Print(err)
		return nil, nil
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		return nil, nil
	}
	if err != nil {
		log.Print(err)
		return nil, nil
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Print(err)
		return nil, nil
	}
	return b, res
}
