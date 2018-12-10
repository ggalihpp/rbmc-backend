package primary

import jwt "github.com/dgrijalva/jwt-go"

// JwtCustomClaims To parse the jwt
type JwtCustomClaims struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	IsAdmin       bool   `json:"is_admin"`
	IsCoordinator bool   `json:"is_coordinator"`
	Territory     string `json:"territory"`
	jwt.StandardClaims
}
