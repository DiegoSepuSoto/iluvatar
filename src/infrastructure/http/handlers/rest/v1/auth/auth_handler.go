package auth

import (
	"fmt"
	"github.com/labstack/echo"
	"iluvatar/src/application/usecase"
	"iluvatar/src/domain/models/requests"
	"iluvatar/src/infrastructure/http/handlers/rest/middleware"
	"iluvatar/src/shared/utils/jwt"
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
	authGroupWithAuthentication.POST("/refresh-token", h.refreshToken)

	authGroupWithoutAuthentication := e.Group(basePath)
	authGroupWithoutAuthentication.POST("/login", h.login)

	return h
}

func (h *authHandler) validateToken(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "valid token", "valid": true})
}

func (h *authHandler) refreshToken(c echo.Context) error {
	token := c.Request().Header.Get(echo.HeaderAuthorization)
	studentID, err := jwt.Token().GetDataFromToken(token, "student_id")
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": fmt.Sprintf("there was an error getting token information [%s]", err.Error())})
	}
	newToken, err := h.authUseCase.RefreshToken(studentID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": fmt.Sprintf("there was an error creating the new token [%s]", err.Error())})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": newToken})
}

func (h *authHandler) login(c echo.Context) error {
	var loginRequest *requests.StudentLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	student, err := h.authUseCase.Login(loginRequest)
	if err != nil {
		if err.Error() == forbiddenUserError {
			return c.JSON(http.StatusForbidden, echo.Map{"message": err.Error()})
		}
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, student)
}
