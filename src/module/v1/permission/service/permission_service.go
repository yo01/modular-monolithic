package service

import (
	"modular-monolithic/module/v1/permission/dto"
	"modular-monolithic/module/v1/permission/helper"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	permissionRepository "modular-monolithic/module/v1/permission/repository"
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
	fetch, err := s.PermissionRepository.PermissionPostgre.Select()
	if err.Error != nil {
		return resp, err
	}

	return helper.PrepareToPermissionsResponse(fetch), err
}

func (s *PermissionService) Detail(id string) (resp *dto.PermissionResponse, merr merror.Error) {
	fetch, err := s.PermissionRepository.PermissionPostgre.SelectByID(id)
	if err.Error != nil {
		return nil, err
	}

	return helper.PrepareToDetailPermissionResponse(fetch), err
}

func (s *PermissionService) Save(req dto.CreatePermissionRequest) (merr merror.Error) {
	if err := s.PermissionRepository.PermissionPostgre.Insert(req); err.Error != nil {
		return err
	}

	return merr
}

func (s *PermissionService) Edit(req dto.UpdatePermissionRequest, id string) (merr merror.Error) {
	if err := s.PermissionRepository.PermissionPostgre.Update(req, id); err.Error != nil {
		return err
	}

	return merr
}

func (s *PermissionService) Delete(id string) (merr merror.Error) {
	if err := s.PermissionRepository.PermissionPostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}
