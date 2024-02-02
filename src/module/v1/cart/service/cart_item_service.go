package service

import (
	"fmt"
	"modular-monolithic/module/v1/cart/dto"
	"modular-monolithic/module/v1/cart/helper"
	cartRepository "modular-monolithic/module/v1/cart/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/google/uuid"
)

type ICartItemService interface {
	List() (resp []dto.CartItemResponse, merr merror.Error)
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

func (s *CartItemService) List() (resp []dto.CartItemResponse, merr merror.Error) {
	fetch, err := s.CartItemRepository.CartItemPostgre.Select()
	if err.Error != nil {
		return resp, err
	}

	return helper.PrepareToCartItemsResponse(fetch), err
}

func (s *CartItemService) Detail(id string) (resp *dto.CartItemResponse, merr merror.Error) {
	fetch, err := s.CartItemRepository.CartItemPostgre.SelectByID(id)
	if err.Error != nil {
		return nil, err
	} else if fetch.ID == uuid.Nil {
		return nil, merror.Error{
			Code:  404,
			Error: fmt.Errorf("cart with id %v is not found", id),
		}
	}

	return helper.PrepareToDetailCartItemResponse(fetch), err
}

func (s *CartItemService) Save(req dto.CreateCartItemRequest) (merr merror.Error) {
	if err := s.CartItemRepository.CartItemPostgre.Insert(req); err.Error != nil {
		return err
	}

	return merr
}

func (s *CartItemService) Edit(req dto.UpdateCartItemRequest, id, cartID string) (merr merror.Error) {
	fetch, _ := s.CartItemRepository.CartItemPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("cart with id %v is not found", id),
		}
	}

	if err := s.CartItemRepository.CartItemPostgre.Update(req, id, cartID); err.Error != nil {
		return err
	}

	return merr
}

func (s *CartItemService) Delete(id string) (merr merror.Error) {
	if err := s.CartItemRepository.CartItemPostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}
