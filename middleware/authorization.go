package middleware

import (
	"strings"

	"lolipad/boilerplate/constant"
	"lolipad/boilerplate/pkg/util"

	"github.com/dgrijalva/jwt-go/v4" //nolint
	"github.com/labstack/echo/v4"    //nolint
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

// Authorization nodoc
func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		bearer := c.Request().Header.Get("Authorization")

		token, err := ExtractBearer(bearer)
		if err != nil {
			return util.ErrorResponse(c, err, nil)
		}

		claims, err := VerifyToken(c, token)
		if err != nil {
			return util.ErrorResponse(c, err, nil)
		}

		c.Set("id", claims[constant.ClaimsID])
		c.Set("token", token)

		return next(c)
	}
}

// ExtractBearer extracts token from bearer
func ExtractBearer(bearer string) (token string, err error) {
	if bearer == "" {
		err = constant.ErrInvalidOrEmptyToken
		return
	}

	splittedBearer := strings.Split(bearer, " ")
	if len(splittedBearer) != 2 {
		err = constant.ErrInvalidOrEmptyToken
		return
	}

	token = splittedBearer[1]

	return
}

// VerifyToken used to verify token and returns claims
func VerifyToken(c echo.Context, token string) (claims jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(
		token,
		func(t *jwt.Token) (i interface{}, err error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return
			}

			key := viper.GetString("encryption.access_token_passphrase")
			i = []byte(key)

			return
		},
	)
	if err != nil {
		if err.Error() == "signature is invalid" {
			err = constant.ErrInvalidOrEmptyToken
		}
		if err.Error() == "Token is expired" {
			err = constant.ErrTokenIsExpired
		}
		return
	}

	claims = jwtToken.Claims.(jwt.MapClaims)
	return
}

// JWTAccessAdminAuth is
func JWTAccessAdminAuth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("jwt.admin_key")),
		SuccessHandler: func(c echo.Context) {
			token := c.Get("user").(*jwt.Token)
			c.Set("admin-token", token.Raw)
		},
	})
}
