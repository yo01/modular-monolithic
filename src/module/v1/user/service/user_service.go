package service

import (
	"fmt"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/user/dto"
	"modular-monolithic/module/v1/user/helper"
	userRepository "modular-monolithic/module/v1/user/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type IUserService interface {
	List(pagination *model.PageRequest) (resp []dto.UserResponse, merr merror.Error)
	Detail(id string) (resp *dto.UserResponse, merr merror.Error)
	Save(req dto.CreateUserRequest) (merr merror.Error)
	Edit(req dto.UpdateUserRequest, id string) (merr merror.Error)
	Delete(id string) (merr merror.Error)
}

type UserService struct {
	Carrier        *mcarrier.Carrier
	UserRepository userRepository.UserRepository
}

func NewUserService(carrier *mcarrier.Carrier) IUserService {
	userRepository := userRepository.NewRepository(carrier)

	return &UserService{
		Carrier:        carrier,
		UserRepository: userRepository,
	}
}

func (s *UserService) List(pagination *model.PageRequest) (resp []dto.UserResponse, merr merror.Error) {
	fetch, err := s.UserRepository.UserPostgre.Select(pagination)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	return helper.PrepareToUsersResponse(fetch), merr
}

func (s *UserService) Detail(id string) (resp *dto.UserResponse, merr merror.Error) {
	fetch, err := s.UserRepository.UserPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	} else if fetch.ID == uuid.Nil {
		err := fmt.Errorf("user with id %v is not found", id)
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  404,
			Error: err,
		}
	}

	return helper.PrepareToDetailUserResponse(fetch), err
}

func (s *UserService) Save(req dto.CreateUserRequest) (merr merror.Error) {
	// HASH PASSWORD
	hashPassword, err := helper.BycryptPassword(req)
	if err != nil {
		zap.S().Error(err)
		return merror.Error{
			Error: err,
		}
	}

	req.Password = hashPassword

	if err := s.UserRepository.UserPostgre.Insert(req); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *UserService) Edit(req dto.UpdateUserRequest, id string) (merr merror.Error) {
	fetch, _ := s.UserRepository.UserPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("user with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  404,
			Error: err,
		}
	}

	if err := s.UserRepository.UserPostgre.Update(req, id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *UserService) Delete(id string) (merr merror.Error) {
	fetch, _ := s.UserRepository.UserPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("user with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  404,
			Error: err,
		}
	}

	if err := s.UserRepository.UserPostgre.Destroy(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}
