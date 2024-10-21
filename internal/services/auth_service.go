package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/mr-emerald-wolf/brew-backend/database"
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

func NewAuthService(repo *db.Queries) IAuthServices {
	return &AuthService{repo: repo}
}

func (as *AuthService) LoginUser(loginRequest req.AuthRequest) (*res.AuthResponse, error) {
	// Retrieve the user by email
	user, err := as.repo.GetUserByEmail(context.Background(), loginRequest.Email)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("user does not exist: %s", loginRequest.Email)
	} else if err != nil {
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

	// Store Refresh Token
	err = database.RedisClient.Set(user.Email, refreshToken, time.Hour)
	if err != nil {
		return nil, err
	}

	response := res.AuthResponse{RefreshToken: refreshToken}
	return &response, nil
}

func (as *AuthService) RefreshToken(rf req.RefreshRequest) (*res.RefreshResponse, error) {

	token, err := jwt.Parse(rf.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET_KEY")), nil
	})

	// Handle token validation errors
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("jwt error: %s", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("jwt error: could not parse claims")
	}

	// Check Refesh Token in Cache
	_, err = database.RedisClient.Get(claims["sub"].(string))
	if err != nil {
		return nil, fmt.Errorf("refresh token not found")
	}

	// Create accessToken
	accessToken, err := CreateToken(claims["sub"].(string), ACCESS_TOKEN)
	if err != nil {
		return nil, fmt.Errorf("jwt error: %s", err.Error())
	}

	response := res.RefreshResponse{
		AccessToken: accessToken,
	}

	return &response, nil
}

func (as *AuthService) LogoutUser(email string) error {
	err := database.RedisClient.Delete(email)
	if err != nil {
		return err
	}
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

	secret := []byte(os.Getenv("ACCESS_SECRET_KEY"))
	expiry := time.Now().Add(time.Minute * 15).Unix()

	if tokenType != ACCESS_TOKEN {
		secret = []byte(os.Getenv("REFRESH_SECRET_KEY"))
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
