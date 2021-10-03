package auth

func (u *authUseCase) RefreshToken(studentID string) (string, error) {
	userToken, err := u.tokenRepository.GenerateJWTForStudent(studentID)
	if err != nil {
		return "", err
	}

	return userToken, nil
}
