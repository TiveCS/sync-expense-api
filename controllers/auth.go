package controllers

import (
	"net/http"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/repositories"
	usecase "github.com/TiveCS/sync-expense/api/usecase/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	Me(ctx echo.Context) error
}

type authController struct {
	userRepo             repositories.UserRepository
	registerUsecase      usecase.AuthRegisterUsecase
	loginUsecase         usecase.AuthLoginUsecase
	meUsecase            usecase.AuthMeUsecase
	generateTokenUsecase usecase.AuthGenerateTokenUsecase
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
	newUser := ctx.Get("payload").(*entities.NewUser)

	result, err := c.registerUsecase.Execute(newUser)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

// Login implements AuthController.
func (c *authController) Login(ctx echo.Context) error {
	loginUser := ctx.Get("payload").(*entities.LoginUser)

	accessToken, refreshToken, err := c.loginUsecase.Execute(loginUser)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func NewAuthController(userRepo repositories.UserRepository) AuthController {
	genTokenUsecase := usecase.NewAuthGenerateTokenUsecase()

	return &authController{
		userRepo:             userRepo,
		registerUsecase:      usecase.NewAuthRegisterUsecase(userRepo),
		loginUsecase:         usecase.NewAuthLoginUsecase(userRepo, genTokenUsecase),
		meUsecase:            usecase.NewAuthMeUsecase(userRepo),
		generateTokenUsecase: genTokenUsecase,
	}
}
