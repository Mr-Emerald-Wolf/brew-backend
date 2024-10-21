package services

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mr-emerald-wolf/brew-backend/internal/db"
	req "github.com/mr-emerald-wolf/brew-backend/internal/dto/request"
	res "github.com/mr-emerald-wolf/brew-backend/internal/dto/response"
	"golang.org/x/crypto/bcrypt"
)

type TokenType int

const (
	ACCESS_TOKEN = iota
	REFRESH_TOKEN
)

type IAuthServices interface {
	LoginUser(req.AuthRequest) (*res.AuthResponse, error)
	RefreshToken(req.RefreshRequest) (*res.RefreshResponse, error)
	LogoutUser(string) error
}

type AuthService struct {
	repo *db.Queries
}

func NewAuthService(repo *db.Queries) AuthService {
	return AuthService{
		repo: repo,
	}
}

func (as AuthService) LoginUser(loginRequest req.AuthRequest) (*res.AuthResponse, error) {
	// Retrieve the user by email
	user, err := as.repo.GetUserByEmail(context.Background(), loginRequest.Email)
	if err != nil {
		return nil, err
	}

	// Verify the provided password against the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password))
	if err != nil {
		// Passwords don't match
		return nil, err
	}

	// Passwords Match
	refreshToken, err := CreateToken(user.Email, REFRESH_TOKEN)

	if err != nil {
		return nil, err
	}

	// TODO: Store Refresh Token

	response := res.AuthResponse{RefreshToken: refreshToken}
	return &response, nil
}

func (as AuthService) RefreshToken(rf req.RefreshRequest) (*res.RefreshResponse, error) {
	return nil, nil
}

func (as AuthService) LogoutUser(uuid string) error {
	return nil
}

// hashPassword hashes a password and returns the hashed value
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CreateToken(email string, tokenType TokenType) (string, error) {

	secret := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	expiry := time.Now().Add(time.Minute * 15).Unix()

	if tokenType != ACCESS_TOKEN {
		secret = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
		expiry = time.Now().Add(time.Hour).Unix()
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,             // Subject (user identifier)
		"iss": "breq-app",        // Issuer
		"aud": "user",            // Audience (user role)
		"exp": expiry,            // Expiration time
		"iat": time.Now().Unix(), // Issued at
	})

	tokenString, err := claims.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
