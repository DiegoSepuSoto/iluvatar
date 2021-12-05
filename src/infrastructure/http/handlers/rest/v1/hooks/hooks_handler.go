package hooks

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"iluvatar/src/application/usecase"
	"iluvatar/src/domain/models/requests"
	"net/http"
)

const postModel = "post"

const basePath = "/v1/hooks"

type hooksHandler struct {
	notificationUseCase usecase.NotificationUseCase
}

func NewHooksHandler(e *echo.Echo, notificationUseCase usecase.NotificationUseCase) *hooksHandler {
	h := &hooksHandler{notificationUseCase: notificationUseCase}

	hooksGroupWithoutAuthentication := e.Group(basePath)
	hooksGroupWithoutAuthentication.POST("/post", h.newPostHookHandler)

	return h
}

func (h *hooksHandler) newPostHookHandler(c echo.Context) error {
	var hookRequest *requests.HookRequest

	if err := c.Bind(&hookRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	if !isPostModel(hookRequest) {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "the request is not a post model"})
	}

	marshalledHookEntry, err := json.Marshal(hookRequest.Entry)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	var newPostHookRequest *requests.HookPostModelRequest
	if errUnmarshal := json.Unmarshal(marshalledHookEntry, &newPostHookRequest); errUnmarshal != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "the request is not a post model"})
	}

	err = h.notificationUseCase.SendNotificationsToAllUsers(newPostHookRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "the notification were sent successfully"})
}

func isPostModel(hookRequest *requests.HookRequest) bool {
	return hookRequest.Model == postModel
}