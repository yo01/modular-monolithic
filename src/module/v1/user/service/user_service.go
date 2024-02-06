package service

import (
	"fmt"
	permissionRepository "modular-monolithic/module/v1/permission/repository"
	roleRepository "modular-monolithic/module/v1/role/repository"
	"modular-monolithic/module/v1/user/dto"
	"modular-monolithic/module/v1/user/helper"
	userRepository "modular-monolithic/module/v1/user/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

type IUserService interface {
	List(subRouterName string) (resp []dto.UserResponse, merr merror.Error)
	Detail(id, subRouterName string) (resp *dto.UserResponse, merr merror.Error)
	Save(req dto.CreateUserRequest, subRouterName string) (merr merror.Error)
	Edit(req dto.UpdateUserRequest, id, subRouterName string) (merr merror.Error)
	Delete(id, subRouterName string) (merr merror.Error)
}

type UserService struct {
	Carrier              *mcarrier.Carrier
	UserRepository       userRepository.UserRepository
	RoleRepository       roleRepository.RoleRepository
	PermissionRepository permissionRepository.PermissionRepository
}

func NewUserService(carrier *mcarrier.Carrier) IUserService {
	userRepository := userRepository.NewRepository(carrier)
	permissionRepository := permissionRepository.NewRepository(carrier)
	roleRepository := roleRepository.NewRepository(carrier)

	return &UserService{
		Carrier:              carrier,
		UserRepository:       userRepository,
		PermissionRepository: permissionRepository,
		RoleRepository:       roleRepository,
	}
}

func (s *UserService) List(subRouterName string) (resp []dto.UserResponse, merr merror.Error) {
	fetch, err := s.UserRepository.UserPostgre.Select()
	if err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	return helper.PrepareToUsersResponse(fetch), merr
}

func (s *UserService) Detail(id, subRouterName string) (resp *dto.UserResponse, merr merror.Error) {
	fetch, err := s.UserRepository.UserPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	} else if len(fetch) == 0 {
		err := fmt.Errorf("user with id %s is not found", id)
		zap.S().Error(err)
		return resp, merror.Error{
			Code:  404,
			Error: err,
		}
	}

	return helper.PrepareToDetailUserResponse(fetch), err
}

func (s *UserService) Save(req dto.CreateUserRequest, subRouterName string) (merr merror.Error) {
	// CHECK USER WITH REQUEST EMAIL IS EXIST OR NOT
	exist, _ := s.UserRepository.UserPostgre.SelectByEmail(req.Email)
	if exist != nil && exist[0].ID != uuid.Nil {
		err := fmt.Errorf("user with email %v is already exist", req.Email)
		zap.S().Error(err.Error)
		return merror.Error{
			Code:  409,
			Error: err,
		}
	}

	// CHECK ROLE IS EXIST OR NOT
	role, _ := s.RoleRepository.RolePostgre.SelectByID(req.RoleID.String())
	if role.ID == uuid.Nil {
		err := fmt.Errorf("role with id %v is not found", req.RoleID.String())
		zap.S().Error(err.Error)
		return merror.Error{
			Code:  404,
			Error: err,
		}
	}

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

func (s *UserService) Edit(req dto.UpdateUserRequest, id, subRouterName string) (merr merror.Error) {
	_, err := s.UserRepository.UserPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	if err := s.UserRepository.UserPostgre.Update(req, id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *UserService) Delete(id, subRouterName string) (merr merror.Error) {
	_, err := s.UserRepository.UserPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	if err := s.UserRepository.UserPostgre.Destroy(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}
