package main

import (
	"context"
	"fmt"
	"github.com/Hirochon/value-api/pkg/logger"
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

	//mysqlClient, err := initialization.NewMySQLClient()
	//if err != nil {
	//	panic(fmt.Sprintf("failed to create mysql client: %s", err))
	//}
	//defer mysqlClient.Close()
}
