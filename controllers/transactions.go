package controllers

import "github.com/TiveCS/sync-expense/api/repositories"

type TransactionController interface{}

type transactionController struct {
	transactionRepository repositories.TransactionRepository
}

func NewTransactionController(tr repositories.TransactionRepository) TransactionController {
	return &transactionController{
		transactionRepository: tr,
	}
}
