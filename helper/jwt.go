package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username, Password string
	CreatedAt          time.Time
	jwt.RegisteredClaims
}

func CreateTokenJwt(username string, password string) (string, error) {
	claims := &Claims{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
