package dtos

import (
	"davidPardoC/rest/domain"

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
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type JWTClaims struct {
	domain.User
	jwt.RegisteredClaims
}
