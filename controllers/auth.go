package controllers

import (
	"net/http"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/repositories"
	accountsUsecase "github.com/TiveCS/sync-expense/api/usecase/accounts"
	authUsecase "github.com/TiveCS/sync-expense/api/usecase/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	Me(ctx echo.Context) error
}

type authController struct {
	registerUsecase      authUsecase.AuthRegisterUsecase
	loginUsecase         authUsecase.AuthLoginUsecase
	meUsecase            authUsecase.AuthMeUsecase
	generateTokenUsecase authUsecase.AuthGenerateTokenUsecase
}

// Me implements AuthController.
func (c *authController) Me(ctx echo.Context) error {
	userToken := ctx.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*entities.JwtUserClaims)

	user, err := c.meUsecase.Execute(claims)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

// Register implements AuthController.
func (c *authController) Register(ctx echo.Context) error {
	newUser := ctx.Get("payload").(*entities.NewUserDTO)

	result, err := c.registerUsecase.Execute(newUser)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

// Login implements AuthController.
func (c *authController) Login(ctx echo.Context) error {
	loginUser := ctx.Get("payload").(*entities.LoginUserDTO)

	accessToken, refreshToken, err := c.loginUsecase.Execute(loginUser)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func NewAuthController(ur repositories.UserRepository, ar repositories.AccountRepository) AuthController {
	gtu := authUsecase.NewAuthGenerateTokenUsecase()

	return &authController{
		registerUsecase:      authUsecase.NewAuthRegisterUsecase(ur, accountsUsecase.NewAccountCreateUsecase(ar)),
		loginUsecase:         authUsecase.NewAuthLoginUsecase(ur, gtu),
		meUsecase:            authUsecase.NewAuthMeUsecase(ur),
		generateTokenUsecase: gtu,
	}
}
