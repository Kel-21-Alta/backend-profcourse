package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	controller "profcourse/controllers"
	"time"
)

type JwtCustomClaims struct {
	Userid string `json:"userid"`
	Role   int8   `json:"role"`
	jwt.StandardClaims
}

type ConfigJwt struct {
	SecretJwt       string
	ExpiredDuration int
}

type JwtRepository interface {
	Int() middleware.JWTConfig
	ExtractClaims(c echo.Context) (*JwtCustomClaims, error)
	GenrateTokenJWT(userId string, role int8) (string, error)
}

func (configJWT *ConfigJwt) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(configJWT.SecretJwt),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(err error, ctx echo.Context) error {
			return controller.NewResponseError(ctx, controller.FORBIDDIN_USER)
		}),
	}
}

func ExtractClaims(c echo.Context) (*JwtCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims, nil
}

func (configJWT *ConfigJwt) GenrateTokenJWT(userId string, role int8) (string, error) {
	claims := JwtCustomClaims{
		userId,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(configJWT.ExpiredDuration))).Unix(),
		},
	}

	// create token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(configJWT.SecretJwt))

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return token, nil
}
