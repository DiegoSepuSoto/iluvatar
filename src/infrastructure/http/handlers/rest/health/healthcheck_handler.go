package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type healthHandler struct{}

type health struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	AppName     string `json:"app_name"`
	Environment string `json:"environment"`
}

func NewHealthHandler(e *echo.Echo) {
	h := &healthHandler{}

	e.GET("/health", h.HealthCheck)
}

func (h *healthHandler) HealthCheck(c echo.Context) error {
	healthCheck := health{
		Status:      "UP",
		AppName:     "Ilúvatar",
		Version:     os.Getenv("APP_VERSION"),
		Environment: os.Getenv("ENV"),
	}

	return c.JSON(http.StatusOK, healthCheck)
}
