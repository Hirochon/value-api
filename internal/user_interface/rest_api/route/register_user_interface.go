package route

import (
	"github.com/Hirochon/value-api/internal/usecase/register_user"
	"github.com/go-logr/logr"
	"io"
	"net/http"
)

// unmarshal する際のフィールドの頭文字は大文字でなければいけない
type registeringUserStruct struct {
	Name     string `json:"name"`
	Gender   int    `json:"gender"`
	BirthDay string `json:"birthDay"`
}

type registeringUser struct {
	registerUserUsecase register_user.RegisterUserUsecase
	registerUserLogger  logr.Logger
}

type RegisterUserInterface interface {
	RegisterUser() func(w http.ResponseWriter, r *http.Request)
}

func NewRegisterUserInterface(registeringUserUsecase register_user.RegisterUserUsecase, registerUserLogger logr.Logger) RegisterUserInterface {
	return registeringUser{
		registerUserUsecase: registeringUserUsecase,
		registerUserLogger:  registerUserLogger,
	}
}

// RegisterUser は、ユーザーを登録する
func (ru registeringUser) RegisterUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			ru.registerUserLogger.Error(err, "failed to read request body")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		parsedRegisteringUser, err := unmarshalToStruct[registeringUserStruct](data)
		if err != nil {
			ru.registerUserLogger.Error(err, "failed to unmarshal request body")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = ru.registerUserUsecase.RegisterUser(parsedRegisteringUser.Name, parsedRegisteringUser.Gender, parsedRegisteringUser.BirthDay)
		if err != nil {
			ru.registerUserLogger.Error(err, "failed to register user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
