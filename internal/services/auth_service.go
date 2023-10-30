package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/mr-emerald-wolf/brew-backend/internal/domain"
	req "github.com/mr-emerald-wolf/brew-backend/internal/dto/request"
	res "github.com/mr-emerald-wolf/brew-backend/internal/dto/response"
	"github.com/mr-emerald-wolf/brew-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type IAuthServices interface {
	LoginUser(req.AuthRequest) (*res.AuthResponse, error)
	RefreshToken(req.RefreshRequest) (*res.RefreshResponse, error)
	// LogoutUser(string) error
}

type AuthService struct {
	r repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return AuthService{
		r: repo,
	}
}

func (as AuthService) LoginUser(loginRequest req.AuthRequest) (*res.AuthResponse, error) {
	// Retrieve the user by email
	user, err := as.r.FindByEmail(loginRequest.Email)
	if err != nil {
		return nil, err
	}

	// Verify the provided password against the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		// Passwords don't match
		return nil, err
	}

	// Passwords Match
	refreshToken, err := GenerateToken(time.Hour*1, domain.TokenPayload{Id: user.UUID, Role: "USER"}, "SECRET_REFRESH")
	if err != nil {
		return nil, err
	}

	// Store Refresh Token
	user.RefreshToken = refreshToken
	_, err = as.r.Update(user.UUID.String(), *user)
	if err != nil {
		return nil, err
	}

	// Generate Access Token
	accessToken, err := GenerateToken(time.Minute*15, domain.TokenPayload{Id: user.UUID, Role: "USER"}, "SECRET")
	if err != nil {
		return nil, err
	}

	response := res.AuthResponse{RefreshToken: refreshToken, AccessToken: accessToken}
	return &response, nil
}

func (as AuthService) RefreshToken(rf req.RefreshRequest) (*res.RefreshResponse, error) {
	token := rf.RefreshToken
	_, err := as.r.FindByRefresh(token)
	if err != nil {
		return nil, err
	}

	// Validate Token
	payload, err := ValidateToken(token, "SECRET_REFRESH")
	if err != nil {
		return nil, err
	}

	// Generate New Access Token
	accessToken, err := GenerateToken(time.Minute*15, payload, "SECRET")

	if err != nil {
		return nil, err
	}

	response := res.RefreshResponse{
		AccessToken: accessToken,
	}
	return &response, nil
}

// hashPassword hashes a password and returns the hashed value
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func GenerateToken(ttl time.Duration, payload domain.TokenPayload, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = payload.Id
	claims["role"] = payload.Role
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTKey))

	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(token string, signedJWTKey string) (domain.TokenPayload, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return domain.TokenPayload{}, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return domain.TokenPayload{}, fmt.Errorf("invalid token claim")
	}
	uuid, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return domain.TokenPayload{}, fmt.Errorf("could not parse uuid: %w", err)
	}
	res := domain.TokenPayload{
		Id:   uuid,
		Role: claims["role"].(string),
	}
	return res, nil
}
