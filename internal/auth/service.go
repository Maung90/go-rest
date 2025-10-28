package auth

import (
	"errors"
	jwtHelper "go-rest/pkg/jwt"
	"go-rest/internal/user"
	"golang.org/x/crypto/bcrypt"
	jwt "github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	ForgotPassword(input ForgotPasswordInput) (string, error)
	Register(input RegisterInput) (user.User, error)
	Login(input LoginInput) (LoginResponse, error)
	ResetPassword(input ResetPasswordInput) error
	FindByEmail(email string) (user.User, error)
	Refresh(refreshToken string) (string, error)
	Logout(refreshToken string) error
}

type authService struct {
	repository Repository
}

func NewAuthService(repository Repository) AuthService {
	return &authService{repository: repository}
}

func (s *authService) Register(input RegisterInput) (user.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user.User{}, err
	}
	newUser := user.User{
		Name:				 input.Name,
		Email:    input.Email,
		Password: string(passwordHash),
	}
	return s.repository.Register(newUser)
}

func (s *authService) FindByEmail(email string) (user.User, error) {
	return s.repository.FindByEmail(email)
}

func (s *authService) Login(input LoginInput) (LoginResponse, error) {
	var response LoginResponse
	foundUser, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return response, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password))
	if err != nil {
		return response, errors.New("invalid email or password")
	}

	tokenDetails, err := jwtHelper.GenerateTokens(foundUser)
	if err != nil {
		return response, err
	}

	err = s.repository.SaveRefreshToken(foundUser.ID, tokenDetails.TokenUUID, tokenDetails.ExpiresAt)
	if err != nil {
		return response, err
	}

	response.AccessToken = tokenDetails.AccessToken
	response.RefreshToken = tokenDetails.RefreshToken
	return response, nil
}

func (s *authService) Refresh(refreshToken string) (string, error) {

	token, err := jwtHelper.ValidateRefreshToken(refreshToken)
	if err != nil || !token.Valid {
		return "", errors.New("invalid or expired refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("error reading token claims")
	}
	tokenUUID := claims["jti"].(string)
	userID, err := s.repository.FetchRefreshToken(tokenUUID)
	if err != nil {
		return "", errors.New("refresh token not found in database")
	}
	tempUser := user.User{ID: userID} 
	newTokens, err := jwtHelper.GenerateTokens(tempUser)
	if err != nil {
		return "", err
	}

	return newTokens.AccessToken, nil
}

func (s *authService) Logout(refreshToken string) error {
	token, err := jwtHelper.ValidateRefreshToken(refreshToken)
	if err != nil || !token.Valid {
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("error reading token claims")
	}
	tokenUUID := claims["jti"].(string)

	return s.repository.DeleteRefreshToken(tokenUUID)
}

func (s *authService) ForgotPassword(input ForgotPasswordInput) (string, error) {
	foundUser, err := s.repository.FindByNameAndEmail(input.Name, input.Email)
	if err != nil {
		return "", errors.New("user with that username and email not found")
	}

	tokenDetails, err := jwtHelper.GenerateResetToken(foundUser)
	if err != nil {
		return "", err
	}

	err = s.repository.SaveRefreshToken(foundUser.ID, tokenDetails.TokenUUID, tokenDetails.ExpiresAt)
	if err != nil {
		return "", err
	}

	return tokenDetails.AccessToken, nil 
}

func (s *authService) ResetPassword(input ResetPasswordInput) error {
	token, err := jwtHelper.ValidateResetToken(input.Token)
	if err != nil || !token.Valid {
		return errors.New("invalid or expired reset token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("error reading token claims")
	}

	tokenUUID := claims["jti"].(string)
	userID, err := s.repository.FetchRefreshToken(tokenUUID)
	if err != nil {
		return errors.New("reset token has already been used or does not exist")
	}
	
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.repository.UpdatePassword(userID, string(newPasswordHash))
	if err != nil {
		return err
	}

	return s.repository.DeleteRefreshToken(tokenUUID)
}