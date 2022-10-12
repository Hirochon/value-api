/*
	custom_error.go は、エラーを生成するためのパッケージです。
	エラーのステータスコードとメッセージを指定してエラーを生成します。
	また、エラーにメッセージを追加することもできます。

	Why: ただのerror型では、ステータスコードを表現できないため。
*/

package custom_error

import (
	"fmt"
	"net/http"
)

// customError は、エラーのメッセージとステータスコードを保持する
type customError struct {
	statusCode int
	message    string
}

// StatusCode は、エラーのメッセージを返す
func (customErrorReceiver customError) Error() string {
	return customErrorReceiver.message
}

// StatusCode は、エラーのステータスコードを返す
func (customErrorReceiver customError) StatusCode() int {
	return customErrorReceiver.statusCode
}

// AddMessage は、エラーにメッセージを追加する
func (customErrorReceiver customError) AddMessage(message string) *customError {
	return &customError{
		statusCode: customErrorReceiver.statusCode,
		message:    customErrorReceiver.message + message,
	}
}

// NewCustomError は、エラーメッセージとステータスコードを指定してエラーを生成する
func NewCustomError(statusCode int, message string) *customError {
	// 全てのstatusCodeを含んだエラーの配列を作成する
	allStatusCodes := []int{
		// 400番台
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusPaymentRequired,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusMethodNotAllowed,
		http.StatusNotAcceptable,
		http.StatusProxyAuthRequired,
		http.StatusRequestTimeout,
		http.StatusConflict,
		http.StatusGone,
		http.StatusLengthRequired,
		http.StatusPreconditionFailed,
		http.StatusRequestEntityTooLarge,
		http.StatusRequestURITooLong,
		http.StatusUnsupportedMediaType,
		http.StatusRequestedRangeNotSatisfiable,
		http.StatusExpectationFailed,
		http.StatusTeapot,
		http.StatusMisdirectedRequest,
		http.StatusUnprocessableEntity,
		http.StatusLocked,
		http.StatusFailedDependency,
		http.StatusTooEarly,
		http.StatusUpgradeRequired,
		http.StatusPreconditionRequired,
		http.StatusTooManyRequests,
		http.StatusRequestHeaderFieldsTooLarge,
		http.StatusUnavailableForLegalReasons,
		// 500番台
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		http.StatusHTTPVersionNotSupported,
		http.StatusVariantAlsoNegotiates,
		http.StatusInsufficientStorage,
		http.StatusLoopDetected,
		http.StatusNotExtended,
		http.StatusNetworkAuthenticationRequired,
	}

	// statusCodeが規定のものでない場合はStatusInternalServerErrorを返す
	for _, constantStatusCode := range allStatusCodes {
		if constantStatusCode == statusCode {
			return &customError{
				statusCode: statusCode,
				message:    message,
			}
		}
	}

	return &customError{
		statusCode: http.StatusInternalServerError,
		message:    fmt.Sprintf("statusCode is not defined: %s", message),
	}
}
