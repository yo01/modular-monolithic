package service

import (
	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"modular-monolithic/module/v1/role/dto"
	"modular-monolithic/module/v1/role/helper"
	roleRepository "modular-monolithic/module/v1/role/repository"
)

type IRoleService interface {
	List() (resp []dto.RoleResponse, merr merror.Error)
	Detail(id string) (resp *dto.RoleResponse, merr merror.Error)
	Save(req dto.CreateRoleRequest) (merr merror.Error)
	Edit(req dto.UpdateRoleRequest, id string) (merr merror.Error)
	Delete(id string) (merr merror.Error)
}

type RoleService struct {
	Carrier        *mcarrier.Carrier
	RoleRepository roleRepository.RoleRepository
}

func NewRoleService(carrier *mcarrier.Carrier) IRoleService {
	roleRepository := roleRepository.NewRepository(carrier)

	return &RoleService{
		Carrier:        carrier,
		RoleRepository: roleRepository,
	}
}

func (s *RoleService) List() (resp []dto.RoleResponse, merr merror.Error) {
	fetch, err := s.RoleRepository.RolePostgre.Select()
	if err.Error != nil {
		return resp, err
	}

	return helper.PrepareToRolesResponse(fetch), err
}

func (s *RoleService) Detail(id string) (resp *dto.RoleResponse, merr merror.Error) {
	fetch, err := s.RoleRepository.RolePostgre.SelectByID(id)
	if err.Error != nil {
		return nil, err
	}

	return helper.PrepareToDetailRoleResponse(fetch), err
}

func (s *RoleService) Save(req dto.CreateRoleRequest) (merr merror.Error) {
	if err := s.RoleRepository.RolePostgre.Insert(req); err.Error != nil {
		return err
	}

	return merr
}

func (s *RoleService) Edit(req dto.UpdateRoleRequest, id string) (merr merror.Error) {
	if err := s.RoleRepository.RolePostgre.Update(req, id); err.Error != nil {
		return err
	}

	return merr
}

func (s *RoleService) Delete(id string) (merr merror.Error) {
	if err := s.RoleRepository.RolePostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}