package register_user

import (
	"github.com/Hirochon/value-api/internal/domain/register_user"
	"github.com/Hirochon/value-api/internal/domain/repository"
	"github.com/go-logr/logr"
)

type RegisterUserUsecase interface {
	RegisterUser(name string, gender int, birthDay string) error
}

type registerUserUsecase struct {
	registerUserRepository repository.RegisterUserRepository
	registerUserLogger     logr.Logger
}

func NewRegisterUserUsecase(registerUserRepository repository.RegisterUserRepository, registerUserLogger logr.Logger) RegisterUserUsecase {
	return registerUserUsecase{
		registerUserRepository: registerUserRepository,
		registerUserLogger:     registerUserLogger,
	}
}

func (ruu registerUserUsecase) RegisterUser(name string, gender int, birthDay string) error {
	registeringUser, err := register_user.NewRegisteringUser(name, gender, birthDay)
	if err != nil {
		return err
	}
	return ruu.registerUserRepository.RegisterUser(registeringUser)
}
