package claim

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	UserId   uint
	UserName string
	jwt.RegisteredClaims
}
