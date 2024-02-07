package service

import (
	"fmt"
	"net/http"

	"modular-monolithic/constant"
	"modular-monolithic/module/v1/permission/dto"
	"modular-monolithic/module/v1/permission/helper"
	permissionRepository "modular-monolithic/module/v1/permission/repository"
	roleRepository "modular-monolithic/module/v1/role/repository"
	"modular-monolithic/security/middleware"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type IPermissionService interface {
	List() (resp []dto.PermissionResponse, merr merror.Error)
	Detail(id string) (resp *dto.PermissionResponse, merr merror.Error)
	Save(req dto.CreatePermissionRequest) (merr merror.Error)
	Edit(req dto.UpdatePermissionRequest, id string) (merr merror.Error)
	Delete(id string) (merr merror.Error)
}

type PermissionService struct {
	Carrier              *mcarrier.Carrier
	PermissionRepository permissionRepository.PermissionRepository
	RoleRepository       roleRepository.RoleRepository
}

func NewRoleService(carrier *mcarrier.Carrier) IPermissionService {
	permissionRepository := permissionRepository.NewRepository(carrier)
	roleRepository := roleRepository.NewRepository(carrier)

	return &PermissionService{
		Carrier:              carrier,
		PermissionRepository: permissionRepository,
		RoleRepository:       roleRepository,
	}
}

func (s *PermissionService) List() (resp []dto.PermissionResponse, merr merror.Error) {
	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, constant.AccessTypePermission, "", nil); err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	fetch, err := s.PermissionRepository.PermissionPostgre.Select()
	if err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	return helper.PrepareToPermissionsResponse(fetch), err
}

func (s *PermissionService) Detail(id string) (resp *dto.PermissionResponse, merr merror.Error) {
	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, constant.AccessTypePermission, "", nil); err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	fetch, err := s.PermissionRepository.PermissionPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	} else if fetch.ID == uuid.Nil {
		err := fmt.Errorf("permission with id %v is not found", id)
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	return helper.PrepareToDetailPermissionResponse(fetch), err
}

func (s *PermissionService) Save(req dto.CreatePermissionRequest) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, constant.AccessTypePermission, "", nil); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	exist, _ := s.PermissionRepository.PermissionPostgre.SelectByRoleID(req.RoleID)
	if exist != nil && exist.ID != uuid.Nil {
		err := fmt.Errorf("permission with role id %s is already exist", req.RoleID)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusConflict,
			Error: err,
		}
	}

	// CHECK ROLE IS EXIST OR NOT
	role, _ := s.RoleRepository.RolePostgre.SelectByID(req.RoleID)
	if role.ID == uuid.Nil {
		err := fmt.Errorf("role with id %s is not found", req.RoleID)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, constant.AccessTypePermission, "", nil); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	if err := s.PermissionRepository.PermissionPostgre.Insert(req); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *PermissionService) Edit(req dto.UpdatePermissionRequest, id string) (merr merror.Error) {
	// CHECK ROLE IS EXIST OR NOT
	role, _ := s.RoleRepository.RolePostgre.SelectByID(req.RoleID)
	if role.ID == uuid.Nil {
		err := fmt.Errorf("role with id %s is not found", req.RoleID)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, constant.AccessTypePermission, "", nil); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	fetch, _ := s.PermissionRepository.PermissionPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("permission with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	if err := s.PermissionRepository.PermissionPostgre.Update(req, id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *PermissionService) Delete(id string) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, constant.AccessTypePermission, "", nil); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	fetch, _ := s.PermissionRepository.PermissionPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("permission with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	if err := s.PermissionRepository.PermissionPostgre.Destroy(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}
