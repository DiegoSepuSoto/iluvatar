package auth

import (
	"iluvatar/src/domain/models"
	"iluvatar/src/shared/utils"
)

func (u *authUseCase) Login(email, password string) (*models.Student, error) {
	student, accessToken, err := u.authRepository.Login(email, password)
	if err != nil {
		return nil, err
	}

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