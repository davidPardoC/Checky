package dtos

import (
	"github.com/golang-jwt/jwt/v5"
)

type LoginDto struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type SignupDto struct {
	Name     string `json:"name"  binding:"required"`
	LastName string `json:"lastName"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type JWTClaims struct {
	UserWithNoPassword
	jwt.RegisteredClaims
}

type UserWithNoPassword struct {
	Name     string `json:"name"  binding:"required"`
	LastName string `json:"lastName"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
}
