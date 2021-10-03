package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	client "github.com/pzentenoe/httpclient-call-go"
	authUseCase "iluvatar/src/application/usecase/auth"
	"iluvatar/src/infrastructure/http/handlers/rest/health"
	authHandler "iluvatar/src/infrastructure/http/handlers/rest/v1/auth"
	authRepository "iluvatar/src/infrastructure/http/repository/miutem/auth"
	"iluvatar/src/infrastructure/http/repository/miutem/career"
	"iluvatar/src/shared/validations"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	e := echo.New()
	e.Validator = validations.NewCustomValidator(validator.New())
	e.HideBanner = true
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	health.NewHealthHandler(e)

	miUTEMAPIHTTPClient := client.NewHTTPClientCall(os.Getenv("MI_UTEM_API_HOST"), &http.Client{})
	authRepositoryImpl := authRepository.NewLoginMiUTEMRepository(miUTEMAPIHTTPClient)
	careerRepository := career.NewCareerMiUTEMRepository(miUTEMAPIHTTPClient)
	authUseCaseImpl := authUseCase.NewAuthUseCase(authRepositoryImpl, careerRepository)
	_ = authHandler.NewAuthHandler(e, authUseCaseImpl)

	quit := make(chan os.Signal, 1)
	go startServer(e, quit)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	gracefulShutdown(e)
}

func startServer(e *echo.Echo, quit chan os.Signal) {
	log.Print("Starting server")
	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		log.Error(err.Error())
		close(quit)
	}
}

func gracefulShutdown(e *echo.Echo) {
	log.Print("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}