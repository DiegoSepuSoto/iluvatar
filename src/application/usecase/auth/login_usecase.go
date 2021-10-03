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

	return student, nil
}
