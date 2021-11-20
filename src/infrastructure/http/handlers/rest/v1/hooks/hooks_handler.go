package hooks

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
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
	return c.JSON(http.StatusOK, echo.Map{"message-received": getJSONRawBody(c)})
}

func getJSONRawBody(c echo.Context) map[string]interface{}  {
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	encodedJson, _ := json.Marshal(jsonBody)
	fmt.Println(string(encodedJson))
	return jsonBody
}