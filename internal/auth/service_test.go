package auth_test

import (
	"projects/GoLinkStat/internal/auth"
	"projects/GoLinkStat/internal/user"
	"testing"
)

const (
	email = "a@a.ru"
)

type mockUserRepository struct{}

func (repo *mockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.ru",
	}, nil
}
func (repo *mockUserRepository) GetByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	service := auth.NewAuthService(&mockUserRepository{})
	registerEmail, err := service.Register(email, "a", "a")
	if err != nil {
		t.Fatal(err)
	}
	if registerEmail != email {
		t.Fatalf("email %s do not match %s", email, registerEmail)
	}
}
