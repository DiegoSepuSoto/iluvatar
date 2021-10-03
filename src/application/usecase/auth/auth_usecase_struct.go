package auth

import "iluvatar/src/domain/repository"

type authUseCase struct {
	authRepository repository.AuthRepository
	careerRepository repository.CareerRepository
	studentRepository repository.StudentRepository
}

func NewAuthUseCase(repositories  ...interface{}) *authUseCase {
	u := new(authUseCase)
	for _, r := range repositories {
		switch t := r.(type) {
			case repository.AuthRepository:
				u.authRepository = t
			case repository.CareerRepository:
				u.careerRepository = t
			case repository.StudentRepository:
				u.studentRepository = t
		}
	}
	return u
}
