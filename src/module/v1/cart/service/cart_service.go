package service

import (
	"fmt"
	"modular-monolithic/module/v1/cart/dto"
	"modular-monolithic/module/v1/cart/helper"
	cartRepository "modular-monolithic/module/v1/cart/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"
)

type ICartService interface {
	List() (resp []dto.CartResponse, merr merror.Error)
	Detail(id string) (resp *dto.CartResponse, merr merror.Error)
	Save(req dto.CreateCartRequest) (resp *dto.CartResponse, merr merror.Error)
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
		zap.S().Error(err.Error)
		return resp, err
	}

	return helper.PrepareToCartsResponse(fetch), err
}

func (s *CartService) Detail(id string) (resp *dto.CartResponse, merr merror.Error) {
	data, err := s.CartRepository.CartPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	} else if len(data) == 0 {
		err := fmt.Errorf("cart with id %s is not found", id)
		zap.S().Error(err)
		return resp, merror.Error{
			Code:  404,
			Error: err,
		}
	}

	return helper.PrepareToDetailCartResponse(data), err
}

func (s *CartService) Save(req dto.CreateCartRequest) (resp *dto.CartResponse, merr merror.Error) {
	fetch, err := s.CartRepository.CartPostgre.Insert(req)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	}

	// GET DETAIL CART
	cart, err := s.CartRepository.CartPostgre.SelectByID(fetch.ID.String())
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	}

	return helper.PrepareToDetailCartResponse(cart), err
}

func (s *CartService) Edit(req dto.UpdateCartRequest, id string) (merr merror.Error) {
	_, err := s.CartRepository.CartPostgre.SelectOneByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	if err := s.CartRepository.CartPostgre.Update(req, id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *CartService) Delete(id string) (merr merror.Error) {
	_, err := s.CartRepository.CartPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	if err := s.CartRepository.CartPostgre.Destroy(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}
