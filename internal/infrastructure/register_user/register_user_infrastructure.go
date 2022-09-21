package register_user

import (
	"github.com/Hirochon/value-api/internal/domain/register_user"
	"github.com/Hirochon/value-api/internal/domain/repository"
	"github.com/go-logr/logr"
	"github.com/jmoiron/sqlx"
)

type registerUserInfrastructure struct {
	mysqlClient        *sqlx.DB
	registerUserLogger logr.Logger
}

func NewRegisterUserRepository(mysqlClient *sqlx.DB, registerUserLogger logr.Logger) repository.RegisterUserRepository {
	return registerUserInfrastructure{
		mysqlClient:        mysqlClient,
		registerUserLogger: registerUserLogger,
	}
}

func (rui registerUserInfrastructure) RegisterUser(registeringUser register_user.RegisteringUser) error {
	rui.registerUserLogger.Info("Reaching RegisterUserRepository")
	rui.registerUserLogger.Info("Registering user", "name", registeringUser.Name(), "gender", registeringUser.Gender(), "birthDay", registeringUser.BirthDayToTime())
	return nil
}
