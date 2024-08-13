package controllers

import (
	"github.com/TiveCS/sync-expense/api/repositories"
	"github.com/labstack/echo/v4"
)

type AccountController interface {
	CreateAccount(ctx *echo.Context) error
	GetAccount(ctx *echo.Context) error
}

type accountsController struct {
	accountsRepository repositories.AccountRepository
}

// CreateAccount implements AccountController.
func (a *accountsController) CreateAccount(ctx *echo.Context) error {
	panic("unimplemented")
}

// GetAccount implements AccountController.
func (a *accountsController) GetAccount(ctx *echo.Context) error {
	panic("unimplemented")
}

func NewAccountController(ar repositories.AccountRepository) AccountController {
	return &accountsController{
		accountsRepository: ar,
	}
}
