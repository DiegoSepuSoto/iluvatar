package auth

import (
	"iluvatar/src/domain/models"
	"iluvatar/src/domain/models/requests"
	"iluvatar/src/shared/utils"
)

func (u *authUseCase) Login(loginRequest *requests.StudentLoginRequest) (*models.Student, error) {
	student, accessToken, err := u.authRepository.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return nil, err
	}

	student.DeviceID = loginRequest.DeviceID

	if accessToken != utils.AdminAccount {
		studentCareer, err := u.careerRepository.GetCareer(accessToken)
		if err != nil {
			return nil, err
		}

		student.Career = studentCareer
	}

	studentID, err := u.studentRepository.UpsertStudent(student)
	if err != nil {
		return nil, err
	}

	userToken, refreshToken, err := u.getUserTokens(studentID)
	if err != nil {
		return nil, err
	}

	student.Token = userToken
	student.RefreshToken = refreshToken

	return student, nil
}

func (u *authUseCase) getUserTokens(studentID string) (string, string, error) {
	userToken, err := u.tokenRepository.GenerateJWTForStudent(studentID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.tokenRepository.GenerateRefreshToken(studentID)
	if err != nil {
		return "", "", err
	}

	return userToken, refreshToken, nil
}