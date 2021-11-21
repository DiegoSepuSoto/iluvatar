package repository

import "iluvatar/src/domain/models/requests"

type NotificationRepository interface {
	SendNotificationsToUsers(newPostHookRequest *requests.HookPostModelRequest, studentsEligibleForNotification []string) error
}