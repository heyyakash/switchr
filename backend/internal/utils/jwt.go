package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Uid   string
	Type  string
	Email string
	Pid   string
}

var sampleSecretKey = []byte(GetString("SECRET_KEY"))

func GenerateJWT(uid string) (string, string, error) {

	expirationTime := time.Now().Add(1 * time.Hour).Unix()
	refreshTokenExpirationTime := time.Now().Add(7 * 24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
		},
		Uid:  uid,
		Type: "auth",
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(refreshTokenExpirationTime, 0)),
		},
		Uid:  uid,
		Type: "auth",
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

func GenerateJWTWithType(email string, _type string, expirationTime int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
		},
		Email: email,
		Type:  _type,
	})

	signedString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return signedString, nil

}
func GenerateApiJWTWithType(pid string, _type string, expirationTime int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
		},
		Pid:  pid,
		Type: _type,
	})

	signedString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return signedString, nil

}

func DecodeJWT(jwtToken string) (UserClaims, bool, error) {
	var userClaim UserClaims

	// Parse the token with claims
	token, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(t *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})

	if err != nil {
		return UserClaims{}, false, err
	}

	// Check if the token is valid and the claims are properly parsed into UserClaims
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		// Validate the expiration time
		valid := claims.ExpiresAt.After(time.Now())
		return *claims, valid, nil
	}

	// If we reach here, something went wrong
	return UserClaims{}, false, nil
}
