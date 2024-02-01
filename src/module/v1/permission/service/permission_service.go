package service

import (
	"fmt"

	"modular-monolithic/module/v1/permission/dto"
	"modular-monolithic/module/v1/permission/helper"
	permissionRepository "modular-monolithic/module/v1/permission/repository"
	"modular-monolithic/module/v1/permission/validation"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

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
}

func NewRoleService(carrier *mcarrier.Carrier) IPermissionService {
	permissionRepository := permissionRepository.NewRepository(carrier)

	return &PermissionService{
		Carrier:              carrier,
		PermissionRepository: permissionRepository,
	}
}

func (s *PermissionService) List() (resp []dto.PermissionResponse, merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidatePermissionAccess(s.Carrier); err.Error != nil {
		return resp, err
	}

	fetch, err := s.PermissionRepository.PermissionPostgre.Select()
	if err.Error != nil {
		return resp, err
	}

	return helper.PrepareToPermissionsResponse(fetch), err
}

func (s *PermissionService) Detail(id string) (resp *dto.PermissionResponse, merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidatePermissionAccess(s.Carrier); err.Error != nil {
		return resp, err
	}

	fetch, err := s.PermissionRepository.PermissionPostgre.SelectByID(id)
	if err.Error != nil {
		return nil, err
	} else if fetch.ID == uuid.Nil {
		return nil, merror.Error{
			Code:  404,
			Error: fmt.Errorf("permission with id %v is not found", id),
		}
	}

	return helper.PrepareToDetailPermissionResponse(fetch), err
}

func (s *PermissionService) Save(req dto.CreatePermissionRequest) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidatePermissionAccess(s.Carrier); err.Error != nil {
		return err
	}

	if err := s.PermissionRepository.PermissionPostgre.Insert(req); err.Error != nil {
		return err
	}

	return merr
}

func (s *PermissionService) Edit(req dto.UpdatePermissionRequest, id string) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidatePermissionAccess(s.Carrier); err.Error != nil {
		return err
	}

	fetch, _ := s.PermissionRepository.PermissionPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("permission with id %v is not found", id),
		}
	}

	if err := s.PermissionRepository.PermissionPostgre.Update(req, id); err.Error != nil {
		return err
	}

	return merr
}

func (s *PermissionService) Delete(id string) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidatePermissionAccess(s.Carrier); err.Error != nil {
		return err
	}

	fetch, _ := s.PermissionRepository.PermissionPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("permission with id %v is not found", id),
		}
	}

	if err := s.PermissionRepository.PermissionPostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}
