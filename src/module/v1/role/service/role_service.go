package service

import (
	"fmt"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/google/uuid"

	"modular-monolithic/module/v1/role/dto"
	"modular-monolithic/module/v1/role/helper"
	roleRepository "modular-monolithic/module/v1/role/repository"
	"modular-monolithic/module/v1/role/validation"
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
	// VALIDATION ACCESS
	if err := validation.ValidateRoleAccess(s.Carrier); err.Error != nil {
		return resp, err
	}

	fetch, err := s.RoleRepository.RolePostgre.Select()
	if err.Error != nil {
		return resp, err
	}

	return helper.PrepareToRolesResponse(fetch), err
}

func (s *RoleService) Detail(id string) (resp *dto.RoleResponse, merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateRoleAccess(s.Carrier); err.Error != nil {
		return resp, err
	}

	fetch, err := s.RoleRepository.RolePostgre.SelectByID(id)
	if err.Error != nil {
		return nil, err
	} else if fetch.ID == uuid.Nil {
		return nil, merror.Error{
			Code:  404,
			Error: fmt.Errorf("role with id %v is not found", id),
		}
	}

	return helper.PrepareToDetailRoleResponse(fetch), err
}

func (s *RoleService) Save(req dto.CreateRoleRequest) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateRoleAccess(s.Carrier); err.Error != nil {
		return err
	}

	if err := s.RoleRepository.RolePostgre.Insert(req); err.Error != nil {
		return err
	}

	return merr
}

func (s *RoleService) Edit(req dto.UpdateRoleRequest, id string) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateRoleAccess(s.Carrier); err.Error != nil {
		return err
	}

	fetch, _ := s.RoleRepository.RolePostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("role with id %v is not found", id),
		}
	}

	if err := s.RoleRepository.RolePostgre.Update(req, id); err.Error != nil {
		return err
	}

	return merr
}

func (s *RoleService) Delete(id string) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateRoleAccess(s.Carrier); err.Error != nil {
		return err
	}

	fetch, _ := s.RoleRepository.RolePostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("role with id %v is not found", id),
		}
	}

	if err := s.RoleRepository.RolePostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}
