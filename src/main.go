package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	client "github.com/pzentenoe/httpclient-call-go"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	_ "iluvatar/docs"
	authUseCase "iluvatar/src/application/usecase/auth"
	notificationUseCase "iluvatar/src/application/usecase/notification"
	sql "iluvatar/src/infrastructure/db/postgresql"
	"iluvatar/src/infrastructure/db/postgresql/ainulindale"
	"iluvatar/src/infrastructure/db/postgresql/ainulindale/entity"
	"iluvatar/src/infrastructure/http/handlers/rest/health"
	authHandler "iluvatar/src/infrastructure/http/handlers/rest/v1/auth"
	hooksHandler "iluvatar/src/infrastructure/http/handlers/rest/v1/hooks"
	messagingRepository "iluvatar/src/infrastructure/http/repository/firebase/messaging"
	authRepository "iluvatar/src/infrastructure/http/repository/miutem/auth"
	"iluvatar/src/infrastructure/http/repository/miutem/career"
	"iluvatar/src/shared/utils/jwt"
	"iluvatar/src/shared/validations"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// @title Documentación Artefacto API
// @version 1.0
// @description En esta documentación se encuentran los detalles de los endpoints presentes en el artefacto API del proyecto Kümelen
// @contact.name Diego Sepúlveda
// @contact.email diego.sepulvedas@utem.cl
func main() {
	postgresqlConnection := CreatePostgreSQLConnection()
	dbConnection := postgresqlConnection.GetConnection()

	e := echo.New()
	e.Validator = validations.NewCustomValidator(validator.New())
	e.HideBanner = true
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	health.NewHealthHandler(e)

	initDBMigrations(dbConnection)

	tokenGenerator := jwt.Token()
	ainulindaleRepository := ainulindale.NewAinulindalePostgresqlSQLRepository(postgresqlConnection)
	miUTEMAPIHTTPClient := client.NewHTTPClientCall(os.Getenv("MI_UTEM_API_HOST"), &http.Client{})
	authRepositoryImpl := authRepository.NewLoginMiUTEMRepository(miUTEMAPIHTTPClient)
	careerRepository := career.NewCareerMiUTEMRepository(miUTEMAPIHTTPClient)
	authUseCaseImpl := authUseCase.NewAuthUseCase(authRepositoryImpl, careerRepository, ainulindaleRepository, tokenGenerator)
	_ = authHandler.NewAuthHandler(e, authUseCaseImpl)

	messagingRepositoryHTTPClient := client.NewHTTPClientCall(os.Getenv("CLOUD_MESSAGE_API_HOST"), &http.Client{})
	messagingRepositoryImpl := messagingRepository.NewMessagingFirebaseRepository(messagingRepositoryHTTPClient)
	newNotificationUseCaseImpl := notificationUseCase.NewNotificationUseCase(messagingRepositoryImpl, ainulindaleRepository)
	_ = hooksHandler.NewHooksHandler(e, newNotificationUseCaseImpl)
	
	quit := make(chan os.Signal, 1)
	go startServer(e, quit)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	gracefulShutdown(e)
}

func initDBMigrations(connection *gorm.DB) {
	err := connection.AutoMigrate(&entity.StudentEntity{})
	if err != nil {
		log.Error(err)
	}
}

func CreatePostgreSQLConnection() *sql.DBConnection {
	postgreSQLHost := os.Getenv("DB_SERVER")
	postgreSQLPort := os.Getenv("DB_PORT")
	postgreSQLDatabase := os.Getenv("DB_NAME")
	postgreSQLUser := os.Getenv("DB_USER")
	postgreSQLPassword := os.Getenv("DB_PASSWORD")

	port, _ := strconv.Atoi(postgreSQLPort)
	connection := sql.NewPostgreSQLConnection(sql.Config().
		SetSQLDialect(sql.PostgresQL).
		Host(postgreSQLHost).
		Port(port).
		DatabaseName(postgreSQLDatabase).
		User(postgreSQLUser).
		Password(postgreSQLPassword),
	)
	return connection
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