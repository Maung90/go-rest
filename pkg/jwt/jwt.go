package jwt

import (
	"errors"
	"go-rest/internal/user"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUUID    string
	ExpiresAt    time.Time
}

func GenerateTokens(user user.User) (*TokenDetails, error) {
	accessLifespan, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_LIFESPAN"))
	if err != nil {
		return nil, err
	}
	refreshLifespan, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_LIFESPAN"))
	if err != nil {
		return nil, err
	}

	td := &TokenDetails{}
	td.ExpiresAt = time.Now().Add(refreshLifespan)
	td.TokenUUID = uuid.NewString()

	accessClaims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Name,
		"exp":      time.Now().Add(accessLifespan).Unix(),
		"iat":      time.Now().Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	td.AccessToken, err = accessToken.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": user.ID,
		"jti":     td.TokenUUID,
		"exp":     td.ExpiresAt.Unix(),
		"iat":     time.Now().Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	td.RefreshToken, err = refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func ValidateRefreshToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})
}