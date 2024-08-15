package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type userClaims struct {
	jwt.RegisteredClaims
	Email string
}

var sampleSecretKey = []byte(GetString("SECRET_KEY"))

func GenerateJWT(email string) (string, string, error) {

	expirationTime := time.Now().Add(1 * time.Hour).Unix()
	refreshTokenExpirationTime := time.Now().Add(7 * 24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
		},
		Email: email,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(refreshTokenExpirationTime, 0)),
		},
		Email: email,
	})

	signedString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", "", err
	}
	signedRefreshToken, err := refreshToken.SignedString(sampleSecretKey)
	if err != nil {
		return "", "", err
	}
	return signedString, signedRefreshToken, nil

}

func DecodeJWT(jwtToken string) (string, error) {
	var userClaim userClaims
	_, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(t *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})

	if err != nil {
		return "", err
	}
	// log.Print(token.Claims)
	return userClaim.Email, nil
}
