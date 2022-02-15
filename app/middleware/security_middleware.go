package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ms-user-portal/app/config"
	"github.com/ms-user-portal/app/logging"
	"github.com/ms-user-portal/app/models"
)

const HeaderKeyAuthorization = "Authorization"

func Authorization(middleware IToken) gin.HandlerFunc {
	return gin.HandlerFunc(middleware.ValidateToken)
}

type IToken interface {
	GenerateToken(ctx *gin.Context, id int) (*string, *models.Error)
	ValidateToken(ctx *gin.Context)
}

type token struct{}

var Token IToken

func NewMiddleware() IToken {
	return &token{}
}

func (t *token) GenerateToken(ctx *gin.Context, id int) (*string, *models.Error) {
	lw := logging.LogForFunc()

	expirationTime := time.Now().Add(time.Minute * 5).Unix()

	claims := &models.Claims{
		ID:        id,
		ExpiresIn: int(expirationTime),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    "Radisys",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JWTSecretKey))

	if err != nil {
		lw.WithField("error", err.Error()).Error("error while generating token")
		return nil, models.NewError(http.StatusInternalServerError, "error while generating token")
	}

	claims.SetLoginCookies(ctx, tokenString)
	return &tokenString, nil
}

func (t *token) ValidateToken(ctx *gin.Context) {
	lw := logging.LogForFunc()
	authToken := ctx.GetHeader(HeaderKeyAuthorization)

	if authToken != "" {
		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWTSecretKey), nil
		})

		if err != nil {
			if err.Error() == jwt.ErrSignatureInvalid.Error() {
				lw.WithField("error", "invalid JWT signature").Error(err)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token Expired"})
			return
		}

		if !token.Valid {
			lw.Error("invalid token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
	} else {
		lw.Error("unauthorized access")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	ctx.Next()
}
