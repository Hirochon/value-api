package route

import (
	"fmt"
	"net/http"
)

// HealthCheck は、httpサーバーのヘルスチェックを行う
func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Health Check OK!")
	w.WriteHeader(http.StatusOK)
}
