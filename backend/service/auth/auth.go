package auth

import (
	"SmartLocker/config"
	"SmartLocker/e"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/log"
)

var key []byte

func JwtSetup() {
	key = []byte(config.Conf.WebServer.JwtSecret)
}

type Claims struct {
	Username string `json:"username"`
	Role     int    `json:"role"`
	Id       int    `json:"id"`
	jwt.StandardClaims
}

func CheckToken(tokenRaw string) (*Claims, int) {
	if tokenRaw == "" {
		return nil, e.InvalidParams
	}

	token, err := jwt.ParseWithClaims(
		tokenRaw,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { //防止通过修改alg为none绕过加密
				return nil,
					fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})

	/*if err != nil {
		log.WithError(err).Debug("parse claims")
		return nil, e.InternalError
	}*/

	if token.Valid {
		c := token.Claims.(*Claims)
		return c, e.Success
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, e.JWTInvalid
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, e.JWTOutOfTime
		} else {
			log.WithError(err).Warn("Couldn't handle this token")
			return nil, e.InternalError
		}
	} else {
		log.WithError(err).Warn("Couldn't handle this token")
		return nil, e.InternalError
	}

}

func (c *Claims) GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	ss, err := token.SignedString(key)
	if err != nil {
		log.WithError(err).Warn("Couldn't sign the token")
		return ""
	}
	return ss
}
