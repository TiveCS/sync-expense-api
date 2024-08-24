package server

import (
	"net/http"

	"github.com/TiveCS/sync-expense/api/enums"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type appValidator struct {
	validator *validator.Validate
}

type AppValidator interface {
	Validate(i interface{}) error
}

func (v *appValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func NewAppValidator() AppValidator {
	v := validator.New()

	v.RegisterValidation("transaction_category", enums.ValidateTransactionCategory)

	return &appValidator{
		validator: v,
	}
}
