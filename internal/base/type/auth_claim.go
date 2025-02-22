package basetype

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaim struct {
	UserID uint
	RoleID *uint
	jwt.RegisteredClaims
}

func NewAuthClaim(userId uint, roleId *uint) AuthClaim {
	return AuthClaim{
		UserID: userId,
		RoleID: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
