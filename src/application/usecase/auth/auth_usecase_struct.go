package auth

import (
	"iluvatar/src/domain/repository"
	"iluvatar/src/shared/utils/jwt"
)

type authUseCase struct {
	authRepository    repository.AuthRepository
	careerRepository  repository.CareerRepository
	studentRepository repository.StudentRepository
	tokenRepository   jwt.JWT
}

func NewAuthUseCase(repositories ...interface{}) *authUseCase {
	u := new(authUseCase)
	for _, r := range repositories {
		switch t := r.(type) {
		case repository.AuthRepository:
			u.authRepository = t
		case repository.CareerRepository:
			u.careerRepository = t
		case repository.StudentRepository:
			u.studentRepository = t
		case jwt.JWT:
			u.tokenRepository = t
		}
	}
	return u
}
