package auth

import "iluvatar/src/domain/models"

func (u *authUseCase) Login(email, password string) (*models.Student, error) {
	student, accessToken, err := u.authRepository.Login(email, password)
	if err != nil {
		return nil, err
	}

	studentCareer, err := u.careerRepository.GetCareer(accessToken)
	if err != nil {
		return nil, err
	}

	student.Career = studentCareer

	err = u.studentRepository.UpsertStudent(student)
	if err != nil {
		return nil, err
	}

	userToken, refreshToken, err := u.getUserTokens(student.Email)
	if err != nil {
		return nil, err
	}

	student.Token = userToken
	student.RefreshToken = refreshToken

	return student, nil
}

func (u *authUseCase) getUserTokens(email string) (string, string, error) {
	userToken, err := u.tokenRepository.GenerateJWTForStudent(email)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.tokenRepository.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	return userToken, refreshToken, nil
}