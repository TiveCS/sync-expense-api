package controllers

import (
	"net/http"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/repositories"
	usecase "github.com/TiveCS/sync-expense/api/usecase/transactions"
	"github.com/labstack/echo/v4"
)

type TransactionController interface {
	NewTransaction(ctx echo.Context) error
	GetTransactionsByOwnerID(ctx echo.Context) error
	GetTransactionDetailsByID(ctx echo.Context) error
	DeleteTransactionByID(ctx echo.Context) error
	EditTransactionByID(ctx echo.Context) error
}

type transactionController struct {
	createTransactionUsecase     usecase.TransactionCreateUsecase
	deleteTransactionUsecase     usecase.TransactionDeleteUsecase
	editTransactionUsecase       usecase.TransactionEditUsecase
	getManyTransactionsUsecase   usecase.TransactionGetManyUsecase
	getDetailsTransactionUsecase usecase.TransactionGetDetailsUsecase
}

// GetTransactionDetailsByID implements TransactionController.
func (t *transactionController) GetTransactionDetailsByID(ctx echo.Context) error {
	transactionID := ctx.Param("transaction_id")

	transaction, err := t.getDetailsTransactionUsecase.Execute(transactionID)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, transaction)
}

// DeleteTransactionByID implements TransactionController.
func (t *transactionController) DeleteTransactionByID(ctx echo.Context) error {
	transactionID := ctx.Param("transaction_id")

	err := t.deleteTransactionUsecase.Execute(transactionID)

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// EditTransactionByID implements TransactionController.
func (t *transactionController) EditTransactionByID(ctx echo.Context) error {
	dto := ctx.Get("payload").(*entities.EditTransactionDTO)

	err := t.editTransactionUsecase.Execute(dto)

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// GetTransactionsByOwnerID implements TransactionController.
func (t *transactionController) GetTransactionsByOwnerID(ctx echo.Context) error {
	dto := ctx.Get("payload").(*entities.GetTransactionsDTO)

	transactions, err := t.getManyTransactionsUsecase.Execute(dto)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, transactions)
}

// NewTransaction implements TransactionController.
func (t *transactionController) NewTransaction(ctx echo.Context) error {
	dto := ctx.Get("payload").(*entities.NewTransactionDTO)

	transactionID, err := t.createTransactionUsecase.Execute(dto)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"transaction_id": transactionID})
}

func NewTransactionController(tr repositories.TransactionRepository) TransactionController {
	return &transactionController{
		createTransactionUsecase:     usecase.NewTransactionCreateUsecase(tr),
		deleteTransactionUsecase:     usecase.NewTransactionDeleteUsecase(tr),
		editTransactionUsecase:       usecase.NewTransactionEditUsecase(tr),
		getManyTransactionsUsecase:   usecase.NewTransactionGetManyUsecase(tr),
		getDetailsTransactionUsecase: usecase.NewTransactionGetDetailsUsecase(tr),
	}
}
