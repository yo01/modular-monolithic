package service

import (
	"modular-monolithic/module/v1/user/dto"
	"modular-monolithic/module/v1/user/helper"
	userRepository "modular-monolithic/module/v1/user/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
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
	}

	return helper.PrepareToDetailUserResponse(fetch), err
}

func (s *UserService) Save(req dto.CreateUserRequest) (merr merror.Error) {
	if err := s.UserRepository.UserPostgre.Insert(req); err.Error != nil {
		return err
	}

	return merr
}

func (s *UserService) Edit(req dto.UpdateUserRequest, id string) (merr merror.Error) {
	err := s.UserRepository.UserPostgre.Update(req, id)

	if err.Error != nil {
		return err
	}

	return merr
}

func (s *UserService) Delete(id string) (merr merror.Error) {
	if err := s.UserRepository.UserPostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}
