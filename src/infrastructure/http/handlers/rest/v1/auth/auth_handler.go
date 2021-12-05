package auth

import (
	"fmt"
	"github.com/labstack/echo/v4"
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

// @Summary Validar Token
// @Description Valida si el token de autenticación enviado es válido
// @Tags API V1 - Autenticación
// @Accept json
// @Produce json
// @Param Token-Autorización header string true "Bearer token"
// @Success 200 {object} map[string]interface{} "OK"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /v1/auth/validate-token [post]
func (h *authHandler) validateToken(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "valid token", "valid": true})
}

// @Summary Actualización Token
// @Description Para reducir la cantidad de veces que un usuario debe iniciar sesión, se diponibiliza un endpoint para actualizar el token de consultas a los diferentes servicios
// @Tags API V1 - Autenticación
// @Accept json
// @Produce json
// @Param Token-Autorización header string true "Bearer token"
// @Success 200 {object} map[string]interface{} "OK"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /v1/auth/refresh-token [post]
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

// @Summary Inicio de Sesión
// @Description Permite a los estudiantes iniciar sesión con sus credenciales de Pasaporte.UTEM
// @Tags API V1 - Autenticación
// @Accept json
// @Produce json
// @Success 200 {object} models.Student "OK"
// @Failure 400 {object} map[string]interface{} "BadRequest"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/auth/login [post]
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
