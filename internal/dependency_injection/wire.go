//go:build wireinject
// +build wireinject

package dependency_injection

import (
	register_user_infrastructure "github.com/Hirochon/value-api/internal/infrastructure/register_user"
	register_user_usecase "github.com/Hirochon/value-api/internal/usecase/register_user"
	"github.com/Hirochon/value-api/internal/user_interface/rest_api/route"
	"github.com/go-logr/logr"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

func InitRegisterUserInterface(sqlDB *sqlx.DB, registerUserLogger logr.Logger) route.RegisterUserInterface {
	wire.Build(route.NewRegisterUserInterface, register_user_usecase.NewRegisterUserUsecase, register_user_infrastructure.NewRegisterUserRepository)
	return nil
}
