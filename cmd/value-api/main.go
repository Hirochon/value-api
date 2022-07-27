package main

import (
	"context"
	"fmt"
	"github.com/Hirochon/value-api/internal/infrastructure/initialization"
	"github.com/Hirochon/value-api/internal/user_interface/rest_api"
	"github.com/Hirochon/value-api/pkg/logger"
	"os"
)

const (
	address = ":8600"
)

// main は、context.Context を初期化して、アプリケーションを起動する
func main() {
	run(context.Background())
}

// run は、アプリケーションを起動する(ログ起動、SQLサーバー起動、APIサーバー起動)
func run(ctx context.Context) {
	l, err := logger.New()
	if err != nil {
		panic(fmt.Sprintf("failed to create logger: %s", err))
	}
	valueApiLogger := l.WithName("value-api")
	valueApiLogger.Info("starting up value-api")

	mysqlClient, err := initialization.NewMySQLClient()
	if err != nil {
		panic(fmt.Sprintf("failed to create mysql client: %s", err))
	}
	defer func() {
		err := mysqlClient.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close mysql client: %s", err))
		}
	}()

	httpErrChan := make(chan error, 1)
	httpServer := rest_api.NewHttpServer(mysqlClient, valueApiLogger)
	go func() {
		if err := httpServer.SetListener(address).Serve(); err != nil {
			httpErrChan <- err
		}
	}()

	select {
	case err := <-httpErrChan:
		valueApiLogger.Error(err, "failed to serve http server")
		os.Exit(1)
	case <-ctx.Done():
		valueApiLogger.Info("shutting down...")
		os.Exit(0)
	}
}
