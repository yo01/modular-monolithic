package service

import (
	"fmt"

	"modular-monolithic/module/v1/product/dto"
	"modular-monolithic/module/v1/product/helper"
	productRepository "modular-monolithic/module/v1/product/repository"
	"modular-monolithic/module/v1/product/validation"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type IProductService interface {
	List() (resp []dto.ProductResponse, merr merror.Error)
	Detail(id string) (resp *dto.ProductResponse, merr merror.Error)
	Save(req dto.CreateProductRequest) (merr merror.Error)
	Edit(req dto.UpdateProductRequest, id string) (merr merror.Error)
	Delete(id string) (merr merror.Error)
}

type ProductService struct {
	Carrier           *mcarrier.Carrier
	ProductRepository productRepository.ProductRepository
}

func NewProductService(carrier *mcarrier.Carrier) IProductService {
	productRepository := productRepository.NewRepository(carrier)

	return &ProductService{
		Carrier:           carrier,
		ProductRepository: productRepository,
	}
}

func (s *ProductService) List() (resp []dto.ProductResponse, merr merror.Error) {
	fetch, err := s.ProductRepository.ProductPostgre.Select()
	if err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	return helper.PrepareToProductsResponse(fetch), err
}

func (s *ProductService) Detail(id string) (resp *dto.ProductResponse, merr merror.Error) {
	fetch, err := s.ProductRepository.ProductPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	} else if fetch.ID == uuid.Nil {
		err := fmt.Errorf("product with id %v is not found", id)
		zap.S().Error(err.Error)
		return nil, merror.Error{
			Code:  404,
			Error: err,
		}
	}

	return helper.PrepareToDetailProductResponse(fetch), err
}

func (s *ProductService) Save(req dto.CreateProductRequest) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateProductAccess(s.Carrier); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	if err := s.ProductRepository.ProductPostgre.Insert(req); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *ProductService) Edit(req dto.UpdateProductRequest, id string) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateProductAccess(s.Carrier); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	fetch, _ := s.ProductRepository.ProductPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("product with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  404,
			Error: err,
		}
	}

	if err := s.ProductRepository.ProductPostgre.Update(req, id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *ProductService) Delete(id string) (merr merror.Error) {
	// VALIDATION ACCESS
	if err := validation.ValidateProductAccess(s.Carrier); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	fetch, _ := s.ProductRepository.ProductPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		err := fmt.Errorf("product with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  404,
			Error: err,
		}
	}

	if err := s.ProductRepository.ProductPostgre.Destroy(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}
