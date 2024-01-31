package service

import (
	"modular-monolithic/module/v1/transaction/dto"
	"modular-monolithic/module/v1/transaction/helper"
	transactionRepository "modular-monolithic/module/v1/transaction/repository"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
)

type ITransactionService interface {
	List() (resp []dto.TransactionResponse, merr merror.Error)
	Detail(id string) (resp *dto.TransactionResponse, merr merror.Error)
	Save(req dto.CreateTransactionRequest) (merr merror.Error)
	Edit(req dto.UpdateTransactionRequest, id string) (merr merror.Error)
	Delete(id string) (merr merror.Error)
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
	if err := s.TransactionRepository.TransactionPostgre.Update(req, id); err.Error != nil {
		return err
	}

	return merr
}

func (s *TransactionService) Delete(id string) (merr merror.Error) {
	if err := s.TransactionRepository.TransactionPostgre.Destroy(id); err.Error != nil {
		return err
	}

	return merr
}
