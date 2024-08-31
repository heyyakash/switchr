package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Uid   string
	Type  string
	Email string
	Role  int
	Pid   string
}

var sampleSecretKey = []byte(GetString("SECRET_KEY"))

func GenerateJWT(uid string) (string, string, error) {

	expirationTime := time.Now().Add(1 * time.Minute).Unix()
	refreshTokenExpirationTime := time.Now().Add(1 * 24 * time.Hour).Unix()

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
func GenerateJWTWithTypeAndUID(uid string, _type string, expirationTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
		},
		Uid:  uid,
		Type: _type,
	})

	signedString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return signedString, nil

}
func GenerateJWTWithTypeUidAndPid(uid string, pid string, role int, _type string, expirationTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
		},
		Uid:  uid,
		Pid:  pid,
		Type: _type,
		Role: role,
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
	_, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(t *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})
	if err != nil {
		log.Print(err)
		if errors.Is(err, jwt.ErrTokenExpired) {
			return userClaim, false, nil
		}
		return userClaim, false, err
	}

	return userClaim, true, nil
}
