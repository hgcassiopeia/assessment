package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hgcassiopeia/assessment/expenses/drivers"
	"github.com/hgcassiopeia/assessment/expenses/handler"
	"github.com/hgcassiopeia/assessment/expenses/repo"
	"github.com/hgcassiopeia/assessment/expenses/service"
)

func main() {
	dbConn := drivers.ConnectDB()
	defer dbConn.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err := drivers.InitTable(dbConn)
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	expensesRepository := repo.InitRepository(dbConn)
	expenseUseCase := service.Init(expensesRepository)
	httpHandler := handler.HttpHandler{UseCase: expenseUseCase}

	e.POST("/expenses", httpHandler.AddNewExpense)

	go func() {
		serverPort := ":" + os.Getenv("PORT")
		if err := e.Start(serverPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server...")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
