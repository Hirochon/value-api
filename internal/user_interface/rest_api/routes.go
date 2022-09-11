package rest_api

import (
	"github.com/Hirochon/value-api/internal/user_interface/rest_api/route"
	"github.com/gorilla/mux"
	"net/http"
)

// registerRoutes は、httpサーバーに実装するルーティングを登録する
func registerRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/health", route.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/user/register", route.RegisterUser).Methods(http.MethodPost)
	return r
}
