// greq - simple http request library
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package greq

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBody(t *testing.T) {
	data := map[string]interface{}{
		"key":     "value",
		"number":  100,
		"boolean": true,
	}
	ts := newTestResponse([]byte{})
	defer ts.Close()
	r := New(ts.URL, true)
	b, err := r.body(data)
	if err != nil {
		t.Error(err)
	}
	re, err := ioutil.ReadAll(b)
	if err != nil {
		t.Error(err)
	}
	ex := "{\"boolean\":true,\"key\":\"value\",\"number\":100}"
	result := string(re)
	if ex != result {
		t.Errorf("\nExpected = %#v\nResult = %#v\n", ex, result)
	}
}

func TestBodyForm(t *testing.T) {
	data := map[string]interface{}{
		"key":     "value",
		"number":  100,
		"boolean": true,
	}
	ts := newTestResponse([]byte{})
	defer ts.Close()
	r := New(ts.URL, false)
	b, err := r.body(data)
	if err != nil {
		t.Error(err)
	}
	re, err := ioutil.ReadAll(b)
	if err != nil {
		t.Error(err)
	}
	ex := "boolean=true&key=value&number=100"
	result := string(re)
	if ex != result {
		t.Errorf("\nExpected = %#v\nResult = %#v\n", ex, result)
	}
}

func TestGet(t *testing.T) {
	b := []byte(`"test"`)
	ts := newTestResponse(b)
	defer ts.Close()
	r := New(ts.URL, true)
	body, _, err := r.Get("/")
	if err != nil {
		t.Errorf("\nExpected = %v\nResult = %v\n", nil, err)
	}
	ex := fmt.Sprintf("%s", string(b))
	result := fmt.Sprintf("%s", string(body))
	if ex != result {
		t.Errorf("\nExpected = %v\nResult = %v\n", ex, result)
	}
}

func TestPost(t *testing.T) {
	ts := newTestReqResponse()
	defer ts.Close()
	r := New(ts.URL, true)
	body, _, err := r.Post("/", map[string]interface{}{"key": "value"})
	if err != nil {
		t.Errorf("\nExpected = %v\nResult = %v\n", nil, err)
	}
	ex := "{\"key\":\"value\"}"
	result := fmt.Sprintf("%s", string(body))
	if ex != result {
		t.Errorf("\nExpected = %v\nResult = %v\n", ex, result)
	}
}

func newTestResponse(response []byte) *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", string(response))
	}))
	return testServer
}
func newTestReqResponse() *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "%s", string(b))
	}))
	return testServer
}
