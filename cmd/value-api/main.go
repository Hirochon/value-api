package main

import (
	"context"
	"fmt"
)

// main は、context.Context を初期化して、アプリケーションを起動する
func main() {
	run(context.Background())
}

// run は、アプリケーションを起動する(ログ起動、SQLサーバー起動、APIサーバー起動)
func run(ctx context.Context) {
	fmt.Println("Hello, world!")
}
