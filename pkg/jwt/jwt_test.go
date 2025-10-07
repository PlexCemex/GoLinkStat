package jwt_test

import (
	"projects/GoLinkStat/pkg/jwt"
	"testing"
)

func TestJWT(t *testing.T) {
	const (
		secret = "oH4tRCK3hFrCRZdo73necAjDBR1i4zz1"
		email  = "a@a.ru"
	)
	jwtService := jwt.NewJWT(secret)
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, parsedJWT := jwtService.Parse(token)
	if !isValid {
		t.Fatal("token is not valid")
	}
	if parsedJWT.Email != email {
		t.Fatal("parsed JWT email != input email")
	}
}
