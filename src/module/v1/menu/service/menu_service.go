package service

import (
	"fmt"

	"modular-monolithic/module/v1/menu/dto"
	"modular-monolithic/module/v1/menu/helper"
	menuRepository "modular-monolithic/module/v1/menu/repository"
	"modular-monolithic/module/v1/menu/validation"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/google/uuid"
)

type IMenuService interface {
	List() (resp []dto.MenuResponse, merr merror.Error)
	Detail(id string) (resp *dto.MenuResponse, merr merror.Error)
	Save(req dto.CreateMenuRequest) (merr merror.Error)
	Edit(req dto.UpdateMenuRequest, id string) (merr merror.Error)
	Delete(id string) (merr merror.Error)
}

type MenuService struct {
	Carrier        *mcarrier.Carrier
	MenuRepository menuRepository.MenuRepository
}

func NewMenuService(carrier *mcarrier.Carrier) IMenuService {
	menuRepository := menuRepository.NewRepository(carrier)

	return &MenuService{
		Carrier:        carrier,
		MenuRepository: menuRepository,
	}
}

func (s *MenuService) List() (resp []dto.MenuResponse, merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateMenuAccess(s.Carrier); err.Error != nil {
		return resp, err
	}

	fetch, err := s.MenuRepository.MenuPostgre.Select()
	if err.Error != nil {
		return resp, err
	}

	return helper.PrepareToMenusResponse(fetch), err
}

func (s *MenuService) Detail(id string) (resp *dto.MenuResponse, merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateMenuAccess(s.Carrier); err.Error != nil {
		return resp, err
	}

	fetch, err := s.MenuRepository.MenuPostgre.SelectByID(id)
	if err.Error != nil {
		return nil, err
	} else if fetch.ID == uuid.Nil {
		return nil, merror.Error{
			Code:  404,
			Error: fmt.Errorf("menu with id %v is not found", id),
		}
	}

	return helper.PrepareToDetailMenuResponse(fetch), err
}

func (s *MenuService) Save(req dto.CreateMenuRequest) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateMenuAccess(s.Carrier); err.Error != nil {
		return err
	}

	if err := s.MenuRepository.MenuPostgre.Insert(req); err.Error != nil {
		return err
	}

	return merr
}

func (s *MenuService) Edit(req dto.UpdateMenuRequest, id string) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateMenuAccess(s.Carrier); err.Error != nil {
		return err
	}

	fetch, _ := s.MenuRepository.MenuPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("menu with id %v is not found", id),
		}
	}

	if err := s.MenuRepository.MenuPostgre.Update(req, id); err.Error != nil {
		return err
	}

	return merr
}

func (s *MenuService) Delete(id string) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateMenuAccess(s.Carrier); err.Error != nil {
		return err
	}

	fetch, _ := s.MenuRepository.MenuPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("menu with id %v is not found", id),
		}
	}

	if err := s.MenuRepository.MenuPostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}
