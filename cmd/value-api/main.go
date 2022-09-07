package main

import (
	"context"
	"fmt"
	"github.com/Hirochon/value-api/internal/infrastructure/initialization"
	"github.com/Hirochon/value-api/internal/pkg/logger"
	"github.com/Hirochon/value-api/internal/user_interface/rest_api"
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
)

const (
	address = ":8600"
)

// main は、context.Context を初期化して、アプリケーションを起動する
func main() {
	os.Exit(run(context.Background()))
}

// run は、アプリケーションを起動する(ログ起動、SQLサーバー起動、APIサーバー起動)
func run(ctx context.Context) (code int) {
	ctx, stop := signal.NotifyContext(ctx, unix.SIGTERM, unix.SIGINT)
	defer stop()

	l, err := logger.New()
	if err != nil {
		_, ferr := fmt.Fprintf(os.Stderr, "failed to create logger: %s", err)
		if ferr != nil {
			panic(fmt.Sprintf("failed to write log:`%s` original error is:`%s`", ferr, err))
		}
		code = 1
		return
	}

	mysqlClient, err := initialization.NewMySQLClient()
	if err != nil {
		_, ferr := fmt.Fprintf(os.Stderr, "failed to create mysql client: %s", err)
		if ferr != nil {
			// Unhandleable, something went wrong...
			panic(fmt.Sprintf("failed to write log:`%s` original error is:`%s`", ferr, err))
		}
		code = 1
		return
	}
	defer func() {
		err := mysqlClient.Close()
		if err != nil {
			_, ferr := fmt.Fprintf(os.Stderr, "failed to close mysql client: %s", err)
			if ferr != nil {
				// Unhandleable, something went wrong...
				panic(fmt.Sprintf("failed to write log:`%s` original error is:`%s`", ferr, err))
			}
		}
	}()

	httpErrChan := make(chan error, 1)
	httpServer := rest_api.NewHttpServer(mysqlClient, l.WithName("value-api"))
	go func() {
		if err := httpServer.SetListener(address).Serve(); err != nil {
			httpErrChan <- err
		}
	}()

	select {
	case err := <-httpErrChan:
		fmt.Println(err.Error())
		code = 1
		return
	case <-ctx.Done():
		fmt.Println("shutting down...")
		code = 0
		return
	}
}
