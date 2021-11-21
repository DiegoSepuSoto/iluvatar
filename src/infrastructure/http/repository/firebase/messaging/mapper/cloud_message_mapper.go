package mapper

import (
	"fmt"
	"iluvatar/src/domain/models/requests"
	"iluvatar/src/infrastructure/http/repository/firebase/messaging/entity"
)

func MapNewPostHookRequestToCloudMessageEntity(newPostHookRequest *requests.HookPostModelRequest, studentEligibleForNotification string) *entity.CloudMessageEntity {
	return &entity.CloudMessageEntity{
		Notification: &entity.Notification{
			Title: fmt.Sprintf("Nueva publicación de %s", newPostHookRequest.Service.ShortName),
			Body:  newPostHookRequest.Title,
		},
		SendTo: studentEligibleForNotification,
	}
}
