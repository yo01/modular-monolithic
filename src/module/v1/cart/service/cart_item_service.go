package service

import (
	"fmt"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/cart/dto"
	"modular-monolithic/module/v1/cart/helper"
	cartRepository "modular-monolithic/module/v1/cart/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type ICartItemService interface {
	List(pageRequest *model.PageRequest) (resp []dto.CartItemResponse, merr merror.Error)
	Detail(id string) (resp *dto.CartItemResponse, merr merror.Error)
	Save(req dto.CreateCartItemRequest) (merr merror.Error)
	Edit(req dto.UpdateCartItemRequest, id, cartID string) (merr merror.Error)
	Delete(id string) (merr merror.Error)
}

type CartItemService struct {
	Carrier            *mcarrier.Carrier
	CartItemRepository cartRepository.CartRepository
}

func NewCartItemService(carrier *mcarrier.Carrier) ICartItemService {
	cartItemRepository := cartRepository.NewRepository(carrier)

	return &CartItemService{
		Carrier:            carrier,
		CartItemRepository: cartItemRepository,
	}
}

func (s *CartItemService) List(pageRequest *model.PageRequest) (resp []dto.CartItemResponse, merr merror.Error) {
	fetch, err := s.CartItemRepository.CartItemPostgre.Select(pageRequest)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	return helper.PrepareToCartItemsResponse(fetch), err
}

func (s *CartItemService) Detail(id string) (resp *dto.CartItemResponse, merr merror.Error) {
	fetch, err := s.CartItemRepository.CartItemPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	} else if fetch.ID == uuid.Nil {
		err := fmt.Errorf("cart with id %v is not found", id)
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  404,
			Error: err,
		}
	}

	return helper.PrepareToDetailCartItemResponse(fetch), err
}

func (s *CartItemService) Save(req dto.CreateCartItemRequest) (merr merror.Error) {
	if err := s.CartItemRepository.CartItemPostgre.Insert(req); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *CartItemService) Edit(req dto.UpdateCartItemRequest, id, cartID string) (merr merror.Error) {
	fetch, _ := s.CartItemRepository.CartItemPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("cart with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  404,
			Error: err,
		}
	}

	if err := s.CartItemRepository.CartItemPostgre.Update(req, id, cartID); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *CartItemService) Delete(id string) (merr merror.Error) {
	if err := s.CartItemRepository.CartItemPostgre.Destroy(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}
