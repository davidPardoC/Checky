package usecases

import (
	"davidPardoC/rest/auth/dtos"
	"davidPardoC/rest/auth/enums"
	"davidPardoC/rest/common"
	"davidPardoC/rest/domain"
	"davidPardoC/rest/users/repository"

	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCases struct {
	userRepo repository.UserRepository
}

func NewAuthUseCases(userRepo repository.UserRepository) *AuthUseCases {
	return &AuthUseCases{userRepo: userRepo}
}

func (useCase *AuthUseCases) SignUpUser(user domain.User) (int64, error) {
	user.Password, _ = hashPassword(user.Password)
	return useCase.userRepo.InsertNewUser(user)
}

func (useCase *AuthUseCases) LoginUser(loginUser dtos.LoginDto) (*dtos.TokenResponse, *common.CustomError) {
	user, err := useCase.userRepo.GetUserByEmail(loginUser.Email)
	if err != nil {
		return nil, err
	}
	isValidPassword := comparePassword(loginUser.Password, user.Password)
	if !isValidPassword {
		return nil, &common.CustomError{StatusCode: http.StatusUnauthorized, Message: "Bad Credentials"}
	}

	return signJwt(*user), nil
}

func (useCase *AuthUseCases) CreateInitialAdminUser() (bool, error) {
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	adminName := os.Getenv("ADMIN_NAME")
	adminLastName := os.Getenv("ADMIN_LAST_NAME")
	user := domain.User{Email: adminEmail, Password: adminPassword, Name: adminName, LastName: adminLastName, Role: enums.AdminRole}
	_, err := useCase.SignUpUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			fmt.Println("Admin user already registered.")
			return true, nil
		}
		return false, err
	}
	fmt.Println("Admin user created.")
	return false, nil
}

func signJwt(user domain.User) *dtos.TokenResponse {
	noPasswordUser := dtos.UserWithNoPassword{Email: user.Email, Name: user.Email, LastName: user.LastName}
	tokenJwtClaims := dtos.JWTClaims{UserWithNoPassword: noPasswordUser, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour))}}
	refreshTokenJwtClaims := dtos.JWTClaims{UserWithNoPassword: noPasswordUser, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour))}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenJwtClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenJwtClaims)

	mySecretKey := os.Getenv("JWT_SECRET")

	tokenString, _ := token.SignedString([]byte(mySecretKey))
	refreshTokenString, _ := refreshToken.SignedString([]byte(mySecretKey))

	return &dtos.TokenResponse{Token: tokenString, RefreshToken: refreshTokenString}

}

func comparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
