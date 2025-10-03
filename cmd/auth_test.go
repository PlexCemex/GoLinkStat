package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"projects/GoLinkStat/internal/auth"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "1",
	})

	resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil{
		t.Fatal(err)
	}
	if resp.StatusCode != 200{
		t.Fatalf("Exepted %d, got %d", 200, resp.StatusCode)
	}
}
