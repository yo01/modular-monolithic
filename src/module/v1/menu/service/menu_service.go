package service

import (
	"fmt"
	"net/http"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/menu/dto"
	"modular-monolithic/module/v1/menu/helper"
	menuRepository "modular-monolithic/module/v1/menu/repository"
	permissionRepository "modular-monolithic/module/v1/permission/repository"
	"modular-monolithic/security/middleware"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type IMenuService interface {
	List(subRouterName string) (resp []dto.MenuResponse, merr merror.Error)
	Detail(id, subRouterName string) (resp *dto.MenuResponse, merr merror.Error)
	Save(req dto.CreateMenuRequest, subRouterName string) (merr merror.Error)
	Edit(req dto.UpdateMenuRequest, id, subRouterName string) (merr merror.Error)
	Delete(id, subRouterName string) (merr merror.Error)
}

type MenuService struct {
	Carrier              *mcarrier.Carrier
	MenuRepository       menuRepository.MenuRepository
	PermissionRepository permissionRepository.PermissionRepository
}

func NewMenuService(carrier *mcarrier.Carrier) IMenuService {
	menuRepository := menuRepository.NewRepository(carrier)
	permissionRepository := permissionRepository.NewRepository(carrier)

	return &MenuService{
		Carrier:              carrier,
		MenuRepository:       menuRepository,
		PermissionRepository: permissionRepository,
	}
}

func (s *MenuService) List(subRouterName string) (resp []dto.MenuResponse, merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	auth := s.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)

	// GET PERMISSION DATA BY ROLE ID
	permission, err := s.PermissionRepository.PermissionPostgre.SelectByRoleID(auth.User.Role.ID.String())
	if err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, "", "", permission.ListAPI); err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	fetch, err := s.MenuRepository.MenuPostgre.Select()
	if err.Error != nil {
		zap.S().Error(err)
		return resp, err
	}

	return helper.PrepareToMenusResponse(fetch), err
}

func (s *MenuService) Detail(id, subRouterName string) (resp *dto.MenuResponse, merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	auth := s.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)

	// GET PERMISSION DATA BY ROLE ID
	permission, err := s.PermissionRepository.PermissionPostgre.SelectByRoleID(auth.User.Role.ID.String())
	if err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, "", "", permission.ListAPI); err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	fetch, err := s.MenuRepository.MenuPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	} else if fetch.ID == uuid.Nil {
		err := fmt.Errorf("menu with id %v is not found", id)
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	return helper.PrepareToDetailMenuResponse(fetch), err
}

func (s *MenuService) Save(req dto.CreateMenuRequest, subRouterName string) (merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	auth := s.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)

	// GET PERMISSION DATA BY ROLE ID
	permission, err := s.PermissionRepository.PermissionPostgre.SelectByRoleID(auth.User.Role.ID.String())
	if err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, "", "", permission.ListAPI); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	if err := s.MenuRepository.MenuPostgre.Insert(req); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *MenuService) Edit(req dto.UpdateMenuRequest, id, subRouterName string) (merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	auth := s.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)

	// GET PERMISSION DATA BY ROLE ID
	permission, err := s.PermissionRepository.PermissionPostgre.SelectByRoleID(auth.User.Role.ID.String())
	if err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, "", "", permission.ListAPI); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	fetch, _ := s.MenuRepository.MenuPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("menu with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	if err := s.MenuRepository.MenuPostgre.Update(req, id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *MenuService) Delete(id, subRouterName string) (merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	auth := s.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)

	// GET PERMISSION DATA BY ROLE ID
	permission, err := s.PermissionRepository.PermissionPostgre.SelectByRoleID(auth.User.Role.ID.String())
	if err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	// VALIDATION ACCESS
	if err := middleware.ValidateAccess(s.Carrier, "", "", permission.ListAPI); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	fetch, _ := s.MenuRepository.MenuPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("menu with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	if err := s.MenuRepository.MenuPostgre.Destroy(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}
