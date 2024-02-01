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

type ICartService interface {
	List() (resp []dto.CartResponse, merr merror.Error)
	Detail(id string) (resp *dto.CartResponse, merr merror.Error)
	Save(req dto.CreateCartRequest) (merr merror.Error)
	Edit(req dto.UpdateCartRequest, id string) (merr merror.Error)
	Delete(id string) (merr merror.Error)
}

type CartService struct {
	Carrier        *mcarrier.Carrier
	CartRepository cartRepository.CartRepository
}

func NewCartService(carrier *mcarrier.Carrier) ICartService {
	cartRepository := cartRepository.NewRepository(carrier)

	return &CartService{
		Carrier:        carrier,
		CartRepository: cartRepository,
	}
}

func (s *CartService) List() (resp []dto.CartResponse, merr merror.Error) {
	fetch, err := s.CartRepository.CartPostgre.Select()
	if err.Error != nil {
		return resp, err
	}

	return helper.PrepareToCartsResponse(fetch), err
}

func (s *CartService) Detail(id string) (resp *dto.CartResponse, merr merror.Error) {
	fetch, err := s.CartRepository.CartPostgre.SelectByID(id)
	if err.Error != nil {
		return nil, err
	} else if fetch.ID == uuid.Nil {
		return nil, merror.Error{
			Code:  404,
			Error: fmt.Errorf("cart with id %v is not found", id),
		}
	}

	return helper.PrepareToDetailCartResponse(fetch), err
}

func (s *CartService) Save(req dto.CreateCartRequest) (merr merror.Error) {
	if err := s.CartRepository.CartPostgre.Insert(req); err.Error != nil {
		return err
	}

	return merr
}

func (s *CartService) Edit(req dto.UpdateCartRequest, id string) (merr merror.Error) {
	fetch, _ := s.CartRepository.CartPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("cart with id %v is not found", id),
		}
	}

	if err := s.CartRepository.CartPostgre.Update(req, id); err.Error != nil {
		return err
	}

	return merr
}

func (s *CartService) Delete(id string) (merr merror.Error) {
	fetch, _ := s.CartRepository.CartPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("cart with id %v is not found", id),
		}
	}

	if err := s.CartRepository.CartPostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}
