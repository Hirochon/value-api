package repository

import (
	"github.com/Hirochon/value-api/internal/domain/register_user"
)

type RegisterUserRepository interface {
	RegisterUser(registeringUser register_user.RegisteringUser) error
}
