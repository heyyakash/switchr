package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

type userClaims struct {
	jwt.RegisteredClaims
	Email string
}

var sampleSecretKey = []byte(GetString("SECRET_KEY"))

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		Email:            email,
	})

	signedString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return signedString, nil

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
