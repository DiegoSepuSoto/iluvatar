package auth

import (
	"github.com/labstack/echo"
	"iluvatar/src/application/usecase"
	"iluvatar/src/domain/models/requests"
	"iluvatar/src/infrastructure/http/handlers/rest/middleware"
	"net/http"
)

const (
	basePath = "/v1/auth"
	forbiddenUserError = "can't login user, wrong http code from response: 403"
)

type authHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(e *echo.Echo, authUseCase usecase.AuthUseCase) *authHandler {
	h := &authHandler{authUseCase: authUseCase}
	authGroupWithAuthentication := e.Group(basePath)
	authGroupWithAuthentication.Use(middleware.ValidateToken())
	authGroupWithAuthentication.POST("/validate-token", h.validateToken)

	authGroupWithoutAuthentication := e.Group(basePath)
	authGroupWithoutAuthentication.POST("/login", h.login)

	return h
}

func (h *authHandler) validateToken(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "valid token", "valid": true})
}

func (h *authHandler) login(c echo.Context) error {
	var loginRequest *requests.StudentLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	student, err := h.authUseCase.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		if err.Error() == forbiddenUserError {
			return c.JSON(http.StatusForbidden, echo.Map{"message": err.Error()})
		}
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, student)
}
