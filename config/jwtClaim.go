package config

import "gopkg.in/dgrijalva/jwt-go.v3"

type JwtClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
