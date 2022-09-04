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

// HttpServer は、app.goでhttpサーバーを構築する際に保持すべき情報
type HttpServer struct {
	Server   *http.Server
	Listener net.Listener
}

// NewHttpServer は、DIとhttpの実装を登録する
// main.goで NewHttpServer を呼び出すことで、httpサーバーを立てる
func NewHttpServer(mysqlClient *sqlx.DB, valueApiLogger logr.Logger) *HttpServer {
	valueApiLogger.Info("starting up value-api")
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(registerRoutes(router)))

	return &HttpServer{
		Server: &http.Server{
			Handler: router,
		},
	}
}

// SetListener は、httpサーバーを立てるために必要なListenerを設定する
// 切り出している理由は、http clientを使ったテストの観点から
func (httpServer *HttpServer) SetListener(address string) *HttpServer {
	return &HttpServer{
		Server: &http.Server{
			Handler: httpServer.Server.Handler,
			Addr:    address,
		},
	}
}

// Serve は、httpサーバーを立てる
func (httpServer *HttpServer) Serve() error {
	return httpServer.Server.ListenAndServe()
}
