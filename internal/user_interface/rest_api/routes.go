package rest_api

import (
	"github.com/gorilla/mux"
	"net/http"
)

// registerRoutes は、httpサーバーに実装するルーティングを登録する
func registerRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/health", healthCheck).Methods(http.MethodGet)
	return r
}

// healthCheck は、httpサーバーのヘルスチェックを行う
func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
