package service

import (
	"net/http"

	"modular-monolithic/config"
	"modular-monolithic/model"
	"modular-monolithic/module/v1/auth/dto"
	"modular-monolithic/module/v1/auth/helper"
	authRepository "modular-monolithic/module/v1/auth/repository"
	userDTO "modular-monolithic/module/v1/user/dto"
	userHelper "modular-monolithic/module/v1/user/helper"
	userRepository "modular-monolithic/module/v1/user/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
	"git.motiolabs.com/library/motiolibs/mtoken"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignIn(req dto.LoginRequest) (resp *dto.LoginResponse, merr merror.Error)
}

type AuthService struct {
	Carrier        *mcarrier.Carrier
	AuthRepository authRepository.AuthRepository
	UserRepository userRepository.UserRepository
}

func NewAuthService(carrier *mcarrier.Carrier) IAuthService {
	authRepository := authRepository.NewRepository(carrier)
	userRepository := userRepository.NewRepository(carrier)

	return &AuthService{
		Carrier:        carrier,
		AuthRepository: authRepository,
		UserRepository: userRepository,
	}
}

func (s *AuthService) SignIn(req dto.LoginRequest) (resp *dto.LoginResponse, merr merror.Error) {
	// Get user details by email
	user, err := s.UserRepository.UserPostgre.SelectByEmail(req.Email)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	}

	userResponse := userHelper.PrepareToLoginDetailUserResponse(user)

	// Verify password
	if err := helper.VerifyPassword(userResponse.Password, req.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		zap.S().Error(err.Error())
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	// Create tokens
	accessToken, refreshToken, err := CreateToken(userResponse)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	}

	// Build the response
	resp = &dto.LoginResponse{
		ID:           userResponse.ID,
		FullName:     userResponse.FullName,
		Email:        userResponse.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// Include role details if available
	if userResponse.Role != nil {
		resp.Role = userResponse.Role
	}

	return resp, merr
}

func CreateToken(user *userDTO.UserLoginResponse) (accessToken, refreshToken string, err merror.Error) {
	// Get config data
	config := config.Get()

	// Build claims
	claims := model.Claims{
		Authorized: true,
		UserID:     user.ID,
		Email:      user.Email,
		FullName:   user.FullName,
		Role:       user.Role,
	}

	// Generate JWT tokens
	accessToken, refreshToken, err = mtoken.GenerateJWTToken(mtoken.JWTConfig{
		SecretKey:              config.JwtKey,
		AccessTokenExpireTime:  int32(config.JwtExpired),
		RefreshTokenExpireTime: int32(config.JwtRefresh),
	}, claims)

	if err.Error != nil {
		zap.S().Error(err.Error)
		return "", "", err
	}

	return accessToken, refreshToken, err
}
