package notification

import "iluvatar/src/domain/repository"

type notificationUseCase struct {
	notificationRepository repository.NotificationRepository
	studentRepository repository.StudentRepository
}

func NewNotificationUseCase(notificationRepository repository.NotificationRepository, studentRepository repository.StudentRepository) *notificationUseCase {
	return &notificationUseCase{notificationRepository: notificationRepository, studentRepository: studentRepository}
}
