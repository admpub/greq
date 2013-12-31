// greq - simple http request library
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package greq

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	b := []byte(`"test"`)
	ts := newTestResponse(b)
	defer ts.Close()
	Host = ts.URL
	body, _ := Get("/")

	ex := fmt.Sprintf("%s", string(b))
	r := fmt.Sprintf("%s", string(body))
	if ex != r {
		t.Errorf("\nExpected = %v\nResult = %v\n", ex, r)
	}
}

func newTestResponse(response []byte) *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", string(response))
	}))
	return testServer
}
