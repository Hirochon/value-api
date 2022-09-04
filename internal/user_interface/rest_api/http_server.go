/*
http_server.go は、main.goで簡単にhttpサーバーを構築するためのパッケージ
*/

package rest_api

import (
	"github.com/go-logr/logr"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net"
	"net/http"
)

// httpServer は、app.goでhttpサーバーを構築する際に保持すべき情報
type httpServer struct {
	server   *http.Server
	listener net.Listener
}

func initCustomRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(registerRoutes(router)))
	return router
}

// NewHttpServer は、DIとhttpの実装を登録する
// main.goで NewHttpServer を呼び出すことで、httpサーバーを立てる
func NewHttpServer(mysqlClient *sqlx.DB, valueApiLogger logr.Logger) *httpServer {
	valueApiLogger.Info("starting up value-api")
	return &httpServer{
		server: &http.Server{
			Handler: initCustomRouter(),
		},
	}
}

// SetListener は、httpサーバーを立てるために必要なListenerを設定する
// 切り出している理由は、http clientを使ったテストの観点から
func (hs *httpServer) SetListener(address string) *httpServer {
	return &httpServer{
		server: &http.Server{
			Handler: hs.server.Handler,
			Addr:    address,
		},
	}
}

// Serve は、httpサーバーを立てる
func (hs *httpServer) Serve() error {
	return hs.server.ListenAndServe()
}
