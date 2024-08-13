package controllers

import (
	"net/http"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/repositories"
	usecase "github.com/TiveCS/sync-expense/api/usecase/accounts"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AccountController interface {
	GetOwnerAccount(ctx echo.Context) error
	EditAccountByID(ctx echo.Context) error
}

type accountsController struct {
	accountsRepository     repositories.AccountRepository
	accountGetOwnedUsecase usecase.AccountGetOwnedUsecase
	accountEditUsecase     usecase.AccountEditUsecase
}

// GetOwnerAccount implements AccountController.
func (c *accountsController) GetOwnerAccount(ctx echo.Context) error {
	userToken := ctx.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*entities.JwtUserClaims)

	account, err := c.accountGetOwnedUsecase.Execute(claims.Subject)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"id":       account.ID,
		"owner_id": account.OwnerID,
		"name":     account.Name,
	})
}

// EditAccountByID implements AccountController.
func (c *accountsController) EditAccountByID(ctx echo.Context) error {
	dto := ctx.Get("payload").(*entities.EditAccountDTO)

	err := c.accountEditUsecase.Execute(dto)

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func NewAccountController(ar repositories.AccountRepository) AccountController {
	return &accountsController{
		accountsRepository:     ar,
		accountGetOwnedUsecase: usecase.NewAccountGetOwnedUsecase(ar),
		accountEditUsecase:     usecase.NewAccountEditUsecase(ar),
	}
}
