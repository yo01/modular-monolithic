package service

import (
	"fmt"
	"net/http"

	"modular-monolithic/config"
	cartRepository "modular-monolithic/module/v1/cart/repository"
	"modular-monolithic/module/v1/transaction/dto"
	"modular-monolithic/module/v1/transaction/helper"
	transactionRepository "modular-monolithic/module/v1/transaction/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"
)

type ITransactionService interface {
	List() (resp []dto.TransactionResponse, merr merror.Error)
	Detail(id string) (resp *dto.TransactionResponse, merr merror.Error)
	Save(req dto.CreateTransactionRequest) (merr merror.Error)
	Edit(req dto.UpdateTransactionRequest, id string) (merr merror.Error)
	Delete(id string) (merr merror.Error)

	// ADDITIONAL
	Payment(id string) (merr merror.Error)
}

type TransactionService struct {
	Carrier               *mcarrier.Carrier
	TransactionRepository transactionRepository.TransactionRepository
	CartRepository        cartRepository.CartRepository
}

func NewTransactionService(carrier *mcarrier.Carrier) ITransactionService {
	transactionRepository := transactionRepository.NewRepository(carrier)
	cartRepository := cartRepository.NewRepository(carrier)

	return &TransactionService{
		Carrier:               carrier,
		TransactionRepository: transactionRepository,
		CartRepository:        cartRepository,
	}
}

func (s *TransactionService) List() (resp []dto.TransactionResponse, merr merror.Error) {
	fetch, err := s.TransactionRepository.TransactionPostgre.Select()
	if err.Error != nil {
		zap.S().Error(err.Error)
		return resp, err
	}

	return helper.PrepareToTransactionsResponse(fetch), merr
}

func (s *TransactionService) Detail(id string) (resp *dto.TransactionResponse, merr merror.Error) {
	fetch, err := s.TransactionRepository.TransactionPostgre.SelectByID(id)
	if err.Error != nil {
		zap.S().Error(err.Error)
		return nil, err
	} else if len(fetch) == 0 {
		err := fmt.Errorf("transaction with id %v is not found", id)
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	return helper.PrepareToDetailTransactionResponse(fetch), err
}

func (s *TransactionService) Save(req dto.CreateTransactionRequest) (merr merror.Error) {
	fetch, _ := s.CartRepository.CartPostgre.SelectByID(req.CartID)
	if len(fetch) > 0 {
		if fetch[0].IsSuccess {
			err := fmt.Errorf("cart with id %v is already success", req.CartID)
			zap.S().Error(err)
			return merror.Error{
				Code:  http.StatusForbidden,
				Error: err,
			}
		}
	}

	if err := s.TransactionRepository.TransactionPostgre.Insert(req); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	if err := SendEmail(); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	// UPDATE FLAG DATA CART IS DONE IF TRANSACTION IS SUCCESS
	if err := s.CartRepository.CartPostgre.UpdateFlagIsSuccess(true, req.CartID); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *TransactionService) Edit(req dto.UpdateTransactionRequest, id string) (merr merror.Error) {
	fetch, _ := s.TransactionRepository.TransactionPostgre.SelectByID(id)
	if len(fetch) == 0 {
		err := fmt.Errorf("transaction with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	if err := s.TransactionRepository.TransactionPostgre.Update(req, id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

func (s *TransactionService) Delete(id string) (merr merror.Error) {
	fetch, _ := s.TransactionRepository.TransactionPostgre.SelectByID(id)
	if len(fetch) == 0 {
		err := fmt.Errorf("transaction with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	if err := s.TransactionRepository.TransactionPostgre.Destroy(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	return merr
}

// ADDITIONAL

func (s *TransactionService) Payment(id string) (merr merror.Error) {
	fetch, _ := s.TransactionRepository.TransactionPostgre.SelectByID(id)
	if len(fetch) == 0 {
		err := fmt.Errorf("transaction with id %v is not found", id)
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	}

	if err := s.TransactionRepository.TransactionPostgre.Payment(id); err.Error != nil {
		zap.S().Error(err.Error)
		return err
	}

	// GET DATA CONFIG
	config := config.Get()

	// dummy data
	data := dto.Email{
		SMTPServer:     config.SMTPServer,
		SMTPPort:       config.SMTPPort,
		SMTPUsername:   config.SMTPUsername,
		SMTPPassword:   config.SMTPPassword,
		SenderEmail:    "yohaneslie0140@gmail.com",
		RecipientEmail: "yohaneslie0140@gmail.com",
		SubjectEmail:   "testing lagi",
	}

	// Create HTML email body using the invoice template
	emailBody, err := helper.GenerateInvoiceHTML(data)
	if err != nil {
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	data.BodyEmail = emailBody

	// Send email
	if err = helper.SendEmail(data); err != nil {
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	return merr
}
