package service

import (
	"fmt"

	"modular-monolithic/module/v1/user/dto"
	"modular-monolithic/module/v1/user/helper"
	userRepository "modular-monolithic/module/v1/user/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/google/uuid"
)

type IUserService interface {
	List() (resp []dto.UserResponse, merr merror.Error)
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

func (s *UserService) List() (resp []dto.UserResponse, merr merror.Error) {
	fetch, err := s.UserRepository.UserPostgre.Select()
	if err.Error != nil {
		return resp, err
	}

	return helper.PrepareToUsersResponse(fetch), merr
}

func (s *UserService) Detail(id string) (resp *dto.UserResponse, merr merror.Error) {
	fetch, err := s.UserRepository.UserPostgre.SelectByID(id)
	if err.Error != nil {
		return nil, err
	} else if fetch.ID == uuid.Nil {
		return nil, merror.Error{
			Code:  404,
			Error: fmt.Errorf("user with id %v is not found", id),
		}
	}

	return helper.PrepareToDetailUserResponse(fetch), err
}

func (s *UserService) Save(req dto.CreateUserRequest) (merr merror.Error) {
	// HASH PASSWORD
	hashPassword, err := helper.BycryptPassword(req)
	if err != nil {
		return merror.Error{
			Error: err,
		}
	}

	req.Password = hashPassword

	if err := s.UserRepository.UserPostgre.Insert(req); err.Error != nil {
		return err
	}

	return merr
}

func (s *UserService) Edit(req dto.UpdateUserRequest, id string) (merr merror.Error) {
	fetch, _ := s.UserRepository.UserPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("user with id %v is not found", id),
		}
	}

	if err := s.UserRepository.UserPostgre.Update(req, id); err.Error != nil {
		return err
	}

	return merr
}

func (s *UserService) Delete(id string) (merr merror.Error) {
	fetch, _ := s.UserRepository.UserPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("user with id %v is not found", id),
		}
	}

	if err := s.UserRepository.UserPostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}
