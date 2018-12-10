package auth

import jwt "github.com/dgrijalva/jwt-go"

type jwtCustomClaims struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"is_admin"`
	Territory string `json:"territory"`
	jwt.StandardClaims
}
