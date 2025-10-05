package main

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"projects/GoLinkStat/internal/auth"
	"projects/GoLinkStat/internal/user"
	"testing"
)

func initDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "a@a.ru",
		Password: "$2a$10$MaODzBkoEI3HV4Q9a60.DO8DFPvg2wIJBf1VnGC/kx4f.e36EKQw.",
		Name:     "a",
	})
}

func TestLoginSuccess(t *testing.T) {
	// Prepare
	db := initDB()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "1",
	})

	resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Exepted %d, got %d", 200, resp.StatusCode)
	}
	token, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(token, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("Token empty")
	}
}

func TestLoginFail(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "2",
	})

	resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 401 {
		t.Fatalf("Exepted %d, got %d", 401, resp.StatusCode)
	}
	// resp.Body.
}
