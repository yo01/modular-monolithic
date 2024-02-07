package service

import (
	"fmt"
	"net/http"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"

	"modular-monolithic/model"
	permissionRepository "modular-monolithic/module/v1/permission/repository"
	"modular-monolithic/module/v1/role/dto"
	"modular-monolithic/module/v1/role/helper"
	roleRepository "modular-monolithic/module/v1/role/repository"
	"modular-monolithic/security/middleware"
)

type IRoleService interface {
	List(subRouterName string) (resp []dto.RoleResponse, merr merror.Error)
	Detail(id, subRouterName string) (resp *dto.RoleResponse, merr merror.Error)
	Save(req dto.CreateRoleRequest, subRouterName string) (merr merror.Error)
	Edit(req dto.UpdateRoleRequest, id, subRouterName string) (merr merror.Error)
	Delete(id, subRouterName string) (merr merror.Error)
}

type RoleService struct {
	Carrier              *mcarrier.Carrier
	RoleRepository       roleRepository.RoleRepository
	PermissionRepository permissionRepository.PermissionRepository
}

func NewRoleService(carrier *mcarrier.Carrier) IRoleService {
	roleRepository := roleRepository.NewRepository(carrier)
	permissionRepository := permissionRepository.NewRepository(carrier)

	return &RoleService{
		Carrier:              carrier,
		RoleRepository:       roleRepository,
		PermissionRepository: permissionRepository,
	}
}

func (s *RoleService) List(subRouterName string) (resp []dto.RoleResponse, merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	context := s.Carrier.Context.Value(middleware.AuthUserCtxKey)

	if context != nil {
		auth := context.(*model.Auth)

		// GET PERMISSION DATA BY ROLE ID
		permission, err := s.PermissionRepository.PermissionPostgre.SelectByRoleID(auth.User.Role.ID.String())
		if err.Error != nil {
			zap.S().Error(err.Error)
			return resp, err
		}

		// VALIDATION ACCESS
		if err := middleware.ValidateAccess(s.Carrier, "", subRouterName, permission.ListAPI); err.Error != nil {
			zap.S().Error(err.Error)
			return resp, err
		}
	}

	fetch, err := s.RoleRepository.RolePostgre.Select()
	if err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	return helper.PrepareToRolesResponse(fetch), err
}

func (s *RoleService) Detail(id, subRouterName string) (resp *dto.RoleResponse, merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	context := s.Carrier.Context.Value(middleware.AuthUserCtxKey)

	if context != nil {
		auth := context.(*model.Auth)

		// GET PERMISSION DATA BY ROLE ID
		permission, err := s.PermissionRepository.PermissionPostgre.SelectByRoleID(auth.User.Role.ID.String())
		if err.Error != nil {
			zap.S().Error(err.Error)
			return resp, err
		}

		// VALIDATION ACCESS
		if err := middleware.ValidateAccess(s.Carrier, "", subRouterName, permission.ListAPI); err.Error != nil {
			zap.S().Error(err.Error)
			return resp, err
		}
	}

	fetch, err := s.RoleRepository.RolePostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	} else if fetch.ID == uuid.Nil {
		err := fmt.Errorf("role with id %v is not found", id)
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	return helper.PrepareToDetailRoleResponse(fetch), err
}

func (s *RoleService) Save(req dto.CreateRoleRequest, subRouterName string) (merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	context := s.Carrier.Context.Value(middleware.AuthUserCtxKey)

	// CHECK ROLE WITH REQUEST NAME IS EXIST OR NOT
	exist, _ := s.RoleRepository.RolePostgre.SelectByName(req.Name)
	if exist != nil && exist.ID != uuid.Nil {
		err := fmt.Errorf("role with name %v is already exist", req.Name)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusConflict,
			Error: err,
		}
	}

	if context != nil {
		auth := context.(*model.Auth)

		// GET PERMISSION DATA BY ROLE ID
		permission, err := s.PermissionRepository.PermissionPostgre.SelectByRoleID(auth.User.Role.ID.String())
		if err.Error != nil {
			zap.S().Error(err.Error)
			return err
		}

		// VALIDATION ACCESS
		if err := middleware.ValidateAccess(s.Carrier, "", subRouterName, permission.ListAPI); err.Error != nil {
			zap.S().Error(err.Error)
			return err
		}
	}

	if err := s.RoleRepository.RolePostgre.Insert(req); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *RoleService) Edit(req dto.UpdateRoleRequest, id, subRouterName string) (merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	auth := s.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)

	// GET PERMISSION DATA BY ROLE ID
	permission, err := s.PermissionRepository.PermissionPostgre.SelectByRoleID(auth.User.Role.ID.String())
	if err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, "", subRouterName, permission.ListAPI); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	fetch, _ := s.RoleRepository.RolePostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("role with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	if err := s.RoleRepository.RolePostgre.Update(req, id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *RoleService) Delete(id, subRouterName string) (merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	auth := s.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)

	// GET PERMISSION DATA BY ROLE ID
	permission, err := s.PermissionRepository.PermissionPostgre.SelectByRoleID(auth.User.Role.ID.String())
	if err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, "", subRouterName, permission.ListAPI); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	fetch, _ := s.RoleRepository.RolePostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("role with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	if err := s.RoleRepository.RolePostgre.Destroy(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}
