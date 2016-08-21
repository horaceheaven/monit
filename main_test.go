package main

import (
	"net/http/httptest"
	"fmt"
	"testing"
	"net/http"
	"io"
)

var (
	server *httptest.Server
	reader io.Reader
	endpoint string
)

func init() {
	server = httptest.NewServer(NewRouter())
	endpoint = fmt.Sprintf("%s/monitor", server.URL)
}

func TestMonitorHandler(t *testing.T) {
	req, err := http.NewRequest("GET", endpoint, reader)

	if err != nil {
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("expect http %s, but got %s", http.StatusOK, res.StatusCode)
	}

	defer res.Body.Close()

	cache := res.Header.Get("Cache-Control")

	if cache != "no-cache" {
		t.Error("expected no-cache in cache control, but got %s", cache)
	}
}
