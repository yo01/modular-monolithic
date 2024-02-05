package service

import (
	"fmt"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/permission/dto"
	"modular-monolithic/module/v1/permission/helper"
	permissionRepository "modular-monolithic/module/v1/permission/repository"
	"modular-monolithic/module/v1/permission/validation"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type IPermissionService interface {
	List(pagination *model.PageRequest) (resp []dto.PermissionResponse, merr merror.Error)
	Detail(id string) (resp *dto.PermissionResponse, merr merror.Error)
	Save(req dto.CreatePermissionRequest) (merr merror.Error)
	Edit(req dto.UpdatePermissionRequest, id string) (merr merror.Error)
	Delete(id string) (merr merror.Error)
}

type PermissionService struct {
	Carrier              *mcarrier.Carrier
	PermissionRepository permissionRepository.PermissionRepository
}

func NewRoleService(carrier *mcarrier.Carrier) IPermissionService {
	permissionRepository := permissionRepository.NewRepository(carrier)

	return &PermissionService{
		Carrier:              carrier,
		PermissionRepository: permissionRepository,
	}
}

func (s *PermissionService) List(pagination *model.PageRequest) (resp []dto.PermissionResponse, merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidatePermissionAccess(s.Carrier); err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	fetch, err := s.PermissionRepository.PermissionPostgre.Select(pagination)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	return helper.PrepareToPermissionsResponse(fetch), err
}

func (s *PermissionService) Detail(id string) (resp *dto.PermissionResponse, merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidatePermissionAccess(s.Carrier); err.Error != nil {
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
			Code:  404,
			Error: err,
		}
	}

	return helper.PrepareToDetailPermissionResponse(fetch), err
}

func (s *PermissionService) Save(req dto.CreatePermissionRequest) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidatePermissionAccess(s.Carrier); err.Error != nil {
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
	// VALIDATION ACCESS
	if err := validation.ValidatePermissionAccess(s.Carrier); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	fetch, _ := s.PermissionRepository.PermissionPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("permission with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  404,
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
	if err := validation.ValidatePermissionAccess(s.Carrier); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	fetch, _ := s.PermissionRepository.PermissionPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("permission with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  404,
			Error: err,
		}
	}

	if err := s.PermissionRepository.PermissionPostgre.Destroy(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}
