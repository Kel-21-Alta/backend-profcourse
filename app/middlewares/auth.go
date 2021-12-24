package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	controller "profcourse/controllers"
	"time"
)

type JwtCustomClaims struct {
	Userid   string `json:"userid"`
	Role     int8   `json:"role"`
	RoleText string `json:"role_text"`
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

func (configJWT *ConfigJwt) GenrateTokenJWT(userId string, role int8, roleText string) (string, error) {
	claims := JwtCustomClaims{
		Userid:   userId,
		Role:     role,
		RoleText: roleText,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(configJWT.ExpiredDuration))).Unix(),
		},
	}

	// create token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(configJWT.SecretJwt))

	if err != nil {
		return "", err
	}
	return token, nil
}
