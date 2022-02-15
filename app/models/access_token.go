package models

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ms-user-portal/app/config"
)

type Claims struct {
	UserName  string `json:"userName"`
	ExpiresIn int    `json:"expiresIn"`
	jwt.StandardClaims
	ID int `json:"id"`
}

// SetLoginCookies => set login httponly cookies on server
func (c *Claims) SetLoginCookies(ctx *gin.Context, token string) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "Access Token",
		Value:    token,
		Domain:   config.Config.Domain,
		MaxAge:   c.ExpiresIn,
		HttpOnly: true,
	})
}

func (c *Claims) DeleteLoginCookies(ctx *gin.Context, token string) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "Access Token",
		Value:    token,
		Domain:   config.Config.Domain,
		MaxAge:   0,
		HttpOnly: true,
	})
}
