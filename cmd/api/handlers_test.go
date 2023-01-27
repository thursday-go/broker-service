package main

import (
	"io"
	"testing"

	"net/http"
	"net/http/httptest"
)

func TestBroker(t *testing.T) {
	app := Config{}

	ts := httptest.NewServer(http.HandlerFunc(app.Broker))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	_, err = io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
}
