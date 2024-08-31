package utils

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCookie(name string, value string, expires time.Time) *http.Cookie {
	secure, err := strconv.ParseBool(GetString("SECURE_COOKIE"))
	if err != nil {
		panic(err)
	}
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		Value:    value,
		Domain:   GetString("DOMAIN"),
		Expires:  expires,
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return cookie
}

func DeleteCookie(name string) *http.Cookie {
	secure, err := strconv.ParseBool(GetString("SECURE_COOKIE"))
	if err != nil {
		panic(err)
	}
	cookie := &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		Domain:   GetString("DOMAIN"),
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return cookie
}

func DeleteTokens(ctx *gin.Context) {
	token := DeleteCookie("token")
	reftoken := DeleteCookie("refreshtoken")
	http.SetCookie(ctx.Writer, token)
	http.SetCookie(ctx.Writer, reftoken)
}
