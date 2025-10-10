package auth_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/internal/auth"
	"projects/GoLinkStat/internal/user"
	"projects/GoLinkStat/pkg/db"
	"testing"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, errors.New("error init mock db")
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}))
	if err != nil {
		return nil, nil, errors.New("error init gorm")
	}
	userRepository := user.NewUserRepository(&db.Db{
		DB: gormDB,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepository),
	}
	return &handler, mock, nil
}

func TestLoginHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("a@a.ru", "$2a$10$MaODzBkoEI3HV4Q9a60.DO8DFPvg2wIJBf1VnGC/kx4f.e36EKQw.")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
		return
	}
	data, err := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "1",
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("got %v, expected %v", w.Result().StatusCode, http.StatusOK)
	}
}
func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	row := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(row)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	if err != nil {
		t.Fatal(err)
		return
	}
	data, err := json.Marshal(&auth.RegisterRequest{
		Email:    "a@a.ru",
		Password: "1",
		Name:     "a",
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(w, r)
	if w.Result().StatusCode != http.StatusCreated {
		t.Fatalf("got %v, expected %v", w.Result().StatusCode, http.StatusCreated)
	}
}
