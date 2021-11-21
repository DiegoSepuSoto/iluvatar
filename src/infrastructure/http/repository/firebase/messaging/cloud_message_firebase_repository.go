package messaging

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"iluvatar/src/domain/models/requests"
	"iluvatar/src/infrastructure/http/repository/firebase/messaging/entity"
	"iluvatar/src/infrastructure/http/repository/firebase/messaging/mapper"
	"io/ioutil"
	"net/http"
	"os"
)

func (r *messagingFirebaseRepository) SendNotificationsToUsers(newPostHookRequest *requests.HookPostModelRequest, studentsEligibleForNotification []string) error {
	for _, studentEligibleForNotification := range studentsEligibleForNotification {
		cloudMessageBody := mapper.MapNewPostHookRequestToCloudMessageEntity(newPostHookRequest, studentEligibleForNotification)

		headers := http.Header{
			echo.HeaderContentType: []string{echo.MIMEApplicationJSON},
			echo.HeaderAuthorization: []string{fmt.Sprintf("key=%s", os.Getenv("CLOUD_MESSAGE_API_KEY"))},
		}

		response, err := r.httpClientCall.
			Path("/fcm/send").
			Method(http.MethodPost).
			Headers(headers).
			Body(cloudMessageBody).
			Do()

		if response.StatusCode != http.StatusOK {
			log.Error(fmt.Sprintf("can't send notification, wrong http code from server: %d", response.StatusCode))
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Error(err.Error())
		}

		var cloudMessageResponse *entity.CloudMessageResponse
		if errorUnmarshal := json.Unmarshal(data, &cloudMessageResponse); errorUnmarshal != nil {
			log.Error("can't unmarshal server response")
		}

		if !isMessageSent(cloudMessageResponse) {
			log.Error("message couldn't be send")
		}

		_ = response.Body.Close()
	}

	return nil
}

func isMessageSent(cloudMessageResponse *entity.CloudMessageResponse) bool {
	return cloudMessageResponse.Failure == 0
}