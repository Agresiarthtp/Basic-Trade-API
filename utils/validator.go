package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTClaim struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
	jwt.RegisteredClaims
}

func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = fmt.Errorf("couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Before(time.Now().UTC()) {
		err = fmt.Errorf("token expired")
		return
	}
	return
}
