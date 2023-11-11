package utils

import (
	"errors"
	"fmt"
	"github.com/xilanhuaer/http-client/common/claim"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenJWT 生成JWT
func GenJWT(userId int32, userName string) (string, error) {
	claims := claim.UserClaim{
		UserId:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "interface",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("UserSecret")))
	return fmt.Sprintf("Bearer %s", tokenString), err
}

// ParseJWT 解析JWT
func ParseJWT(tokenString string) (*claim.UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claim.UserClaim{}, security())
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*claim.UserClaim); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("couldn't handle this token")
	}
}

func security() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("UserSecret")), nil
	}
}
