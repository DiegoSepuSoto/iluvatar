package usecase

import "iluvatar/src/domain/models/requests"

type NotificationUseCase interface {
	SendNotificationsToAllUsers(newPostHookRequest *requests.HookPostModelRequest) error
}