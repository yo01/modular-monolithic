package service

import (
	"modular-monolithic/config"
	"modular-monolithic/module/v1/auth/dto"
	"modular-monolithic/module/v1/auth/helper"
	authRepository "modular-monolithic/module/v1/auth/repository"
	userRepository "modular-monolithic/module/v1/user/repository"

	"time"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/golang-jwt/jwt"
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
	// GET DETAIL USER BY EMAIL
	user, err := s.UserRepository.UserPostgre.SelectByEmail(req.Email)
	if err.Error != nil {
		return nil, err
	}

	if err := helper.VerifyPassword(*user.Password, req.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	token, error := CreateToken(user.ID.String())
	if error != nil {
		return nil, merror.Error{
			Code:  500,
			Error: error,
		}
	}

	resp = new(dto.LoginResponse)
	resp.ID = user.ID
	resp.FullName = user.FullName
	resp.Email = user.Email
	resp.Token = token

	return resp, merr
}

func CreateToken(user_id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Local().Add(time.Hour * 24 * 7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// GET CONFIG DATA
	config := config.Get()

	return token.SignedString([]byte(config.AppApiKey))
}
