package service

import (
	"fmt"

	"modular-monolithic/config"
	"modular-monolithic/module/v1/transaction/dto"
	"modular-monolithic/module/v1/transaction/helper"
	transactionRepository "modular-monolithic/module/v1/transaction/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/google/uuid"
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
}

func NewTransactionService(carrier *mcarrier.Carrier) ITransactionService {
	transactionRepository := transactionRepository.NewRepository(carrier)

	return &TransactionService{
		Carrier:               carrier,
		TransactionRepository: transactionRepository,
	}
}

func (s *TransactionService) List() (resp []dto.TransactionResponse, merr merror.Error) {
	fetch, err := s.TransactionRepository.TransactionPostgre.Select()
	if err.Error != nil {
		return resp, err
	}

	return helper.PrepareToTransactionsResponse(fetch), merr
}

func (s *TransactionService) Detail(id string) (resp *dto.TransactionResponse, merr merror.Error) {
	fetch, err := s.TransactionRepository.TransactionPostgre.SelectByID(id)
	if err.Error != nil {
		return nil, err
	} else if fetch.ID == uuid.Nil {
		return nil, merror.Error{
			Code:  404,
			Error: fmt.Errorf("transaction with id %v is not found", id),
		}
	}

	return helper.PrepareToDetailTransactionResponse(fetch), err
}

func (s *TransactionService) Save(req dto.CreateTransactionRequest) (merr merror.Error) {
	if err := s.TransactionRepository.TransactionPostgre.Insert(req); err.Error != nil {
		return err
	}

	return merr
}

func (s *TransactionService) Edit(req dto.UpdateTransactionRequest, id string) (merr merror.Error) {
	fetch, _ := s.TransactionRepository.TransactionPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("transaction with id %v is not found", id),
		}
	}

	if err := s.TransactionRepository.TransactionPostgre.Update(req, id); err.Error != nil {
		return err
	}

	return merr
}

func (s *TransactionService) Delete(id string) (merr merror.Error) {
	fetch, _ := s.TransactionRepository.TransactionPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("transaction with id %v is not found", id),
		}
	}

	if err := s.TransactionRepository.TransactionPostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}

// ADDITIONAL

func (s *TransactionService) Payment(id string) (merr merror.Error) {
	fetch, _ := s.TransactionRepository.TransactionPostgre.SelectByID(id)
	if fetch.ID == uuid.Nil {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("transaction with id %v is not found", id),
		}
	}

	if err := s.TransactionRepository.TransactionPostgre.Payment(id); err.Error != nil {
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
		return merror.Error{
			Code:  500,
			Error: err,
		}
	}

	data.BodyEmail = emailBody

	// Send email
	if err = helper.SendEmail(data); err != nil {
		return merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return merr
}
