package hooks

import (
	"github.com/labstack/echo"
	"net/http"
)

const basePath = "/v1/hooks"

type hooksHandler struct {}

func NewHooksHandler(e *echo.Echo) *hooksHandler {
	h := &hooksHandler{}

	hooksGroupWithoutAuthentication := e.Group(basePath)
	hooksGroupWithoutAuthentication.POST("/post", h.newPostHookHandler)

	return h
}

func (h *hooksHandler) newPostHookHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message-received": c.Request().Body})
}