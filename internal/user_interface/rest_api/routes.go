package rest_api

import (
	"github.com/Hirochon/value-api/internal/dependency_injection"
	"github.com/Hirochon/value-api/internal/user_interface/rest_api/route"
	"github.com/go-logr/logr"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
)

// registerRoutes は、httpサーバーに実装するルーティングを登録する
func registerRoutes(r *mux.Router, mysqlClient *sqlx.DB, valueApiLogger logr.Logger) *mux.Router {
	registerUserInterface := dependency_injection.InitRegisterUserInterface(mysqlClient, valueApiLogger)
	r.HandleFunc("/health", route.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/user/register", registerUserInterface.RegisterUser()).Methods(http.MethodPost)
	return r
}
