package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	client "github.com/pzentenoe/httpclient-call-go"
	"gorm.io/gorm"
	authUseCase "iluvatar/src/application/usecase/auth"
	sql "iluvatar/src/infrastructure/db/postgresql"
	"iluvatar/src/infrastructure/db/postgresql/ainulindale"
	"iluvatar/src/infrastructure/db/postgresql/ainulindale/entity"
	"iluvatar/src/infrastructure/http/handlers/rest/health"
	authHandler "iluvatar/src/infrastructure/http/handlers/rest/v1/auth"
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

func main() {
	postgresqlConnection := CreatePostgreSQLConnection()
	dbConnection := postgresqlConnection.GetConnection()

	e := echo.New()
	e.Validator = validations.NewCustomValidator(validator.New())
	e.HideBanner = true
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	health.NewHealthHandler(e)

	initDBMigrations(dbConnection)

	tokenGenerator := jwt.Token()
	ainulindaleRepository := ainulindale.NewAinulindalePostgresqlSQLRepository(postgresqlConnection)
	miUTEMAPIHTTPClient := client.NewHTTPClientCall(os.Getenv("MI_UTEM_API_HOST"), &http.Client{})
	authRepositoryImpl := authRepository.NewLoginMiUTEMRepository(miUTEMAPIHTTPClient)
	careerRepository := career.NewCareerMiUTEMRepository(miUTEMAPIHTTPClient)
	authUseCaseImpl := authUseCase.NewAuthUseCase(authRepositoryImpl, careerRepository, ainulindaleRepository, tokenGenerator)
	_ = authHandler.NewAuthHandler(e, authUseCaseImpl)

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