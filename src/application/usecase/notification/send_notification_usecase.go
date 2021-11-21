package notification

import (
	"errors"
	"iluvatar/src/domain/models/requests"
)

func (u *notificationUseCase) SendNotificationsToAllUsers(newPostHookRequest *requests.HookPostModelRequest) error {
	studentsEligibleForNotification, err := u.studentRepository.GetStudentsEligibleForNotification()
	if err != nil {
		return err
	}

	if len(studentsEligibleForNotification) == 0 {
		return errors.New("no students eligible for notification were found")
	}

	err = u.notificationRepository.SendNotificationsToUsers(newPostHookRequest, studentsEligibleForNotification)
	if err != nil {
		return err
	}

	return nil
}
