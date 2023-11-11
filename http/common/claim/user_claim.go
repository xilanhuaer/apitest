package claim

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	UserId   int32
	UserName string
	jwt.RegisteredClaims
}
