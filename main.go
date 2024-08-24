package main

import (
	"github.com/TiveCS/sync-expense/api/controllers"
	"github.com/TiveCS/sync-expense/api/db"
	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/middlewares"
	"github.com/TiveCS/sync-expense/api/repositories"
	"github.com/TiveCS/sync-expense/api/server"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Validator = server.NewAppValidator()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	db := db.Connect()

	userRepo := repositories.NewUserRepository(db)
	accountRepo := repositories.NewAccountRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	authController := controllers.NewAuthController(userRepo, accountRepo)
	accountsController := controllers.NewAccountController(accountRepo)
	transactionsController := controllers.NewTransactionController(transactionRepo)

	v1 := e.Group("/api/v1")

	auth := v1.Group("/auth")
	auth.POST("/register", authController.Register, middlewares.Validate(&entities.NewUserDTO{}))
	auth.POST("/login", authController.Login, middlewares.Validate(&entities.LoginUserDTO{}))
	auth.GET("/me", authController.Me, middlewares.Authenticated())

	accounts := v1.Group("/accounts", middlewares.Authenticated())
	accounts.GET("", accountsController.GetOwnerAccount, middlewares.Validate(&entities.GetManyAccountsDTO{}))
	accounts.PUT("/:account_id", accountsController.EditAccountByID, middlewares.Validate(&entities.EditAccountDTO{}))

	transactions := v1.Group("/transactions", middlewares.Authenticated())
	transactions.POST("", transactionsController.NewTransaction, middlewares.Validate(&entities.NewTransactionDTO{}))
	transactions.GET("", transactionsController.GetTransactionsByOwnerID, middlewares.Validate(&entities.GetTransactionsDTO{}))
	transactions.GET("/:transaction_id", transactionsController.GetTransactionDetailsByID)
	transactions.DELETE("/:transaction_id", transactionsController.DeleteTransactionByID)
	transactions.PUT("/:transaction_id", transactionsController.EditTransactionByID, middlewares.Validate(&entities.EditTransactionDTO{}))

	e.Logger.Fatal(e.Start(":1323"))
}
