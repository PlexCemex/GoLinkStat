package jwt

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret: secret,
	}
}

func (j *JWT) Create(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"Email": email,
	})
	s, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return s, nil
}
