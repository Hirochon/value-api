package custom_error

import (
	"net/http"
	"testing"
)

// TestCustomError は、custom_errorパッケージの正常系テストを行う
func TestCustomErrorSuccess(t *testing.T) {
	t.Parallel()

	cases := []struct {
		scenario       string
		errMessage     string
		addErrMessage  string
		errStatusCode  int
		wantErrMessage string
		wantStatusCode int
	}{
		{
			scenario:       "正常系:エラーメッセージとステータスコードを指定してエラーを生成",
			errMessage:     "エラーメッセージ",
			errStatusCode:  http.StatusNotFound,
			wantErrMessage: "エラーメッセージ",
			wantStatusCode: http.StatusNotFound,
		},
		{
			scenario:       "正常系:エラーメッセージとステータスコードを指定してエラーを生成",
			errMessage:     "ほげりんちょ",
			errStatusCode:  http.StatusInternalServerError,
			wantErrMessage: "ほげりんちょ",
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			scenario:       "正常系:エラーメッセージとステータスコードを指定してエラーを生成",
			errMessage:     "はげりんちょ",
			addErrMessage:  "つるりんちょ",
			errStatusCode:  http.StatusRequestedRangeNotSatisfiable,
			wantErrMessage: "はげりんちょつるりんちょ",
			wantStatusCode: http.StatusRequestedRangeNotSatisfiable,
		},
	}

	for _, c := range cases {
		t.Run(c.scenario, func(t *testing.T) {
			err := NewCustomError(c.errStatusCode, c.errMessage)
			err = err.AddMessage(c.addErrMessage)
			if err == nil {
				t.Errorf("エラーが生成されていません。")
			}
			if err.Error() != c.wantErrMessage {
				t.Errorf("期待していたエラーメッセージと生成したエラーメッセージが一致しません。 want:%s, receive:%s", c.wantErrMessage, err.Error())
			}
			if err.StatusCode() != c.wantStatusCode {
				t.Errorf("期待していたステータスコードと生成したステータスコードが一致しません。 want:%d, receive:%d", c.wantStatusCode, err.StatusCode())
			}
		})
	}
}

// TestCustomErrorFail は、custom_errorパッケージの異常系テストを行う
func TestCustomErrorFail(t *testing.T) {
	t.Parallel()
	cases := []struct {
		scenario      string
		errMessage    string
		errStatusCode int
	}{
		{
			scenario:      "異常系:ステータスコードが0の場合インターナルサーバーエラーを生成",
			errMessage:    "エラーメッセージ",
			errStatusCode: 0,
		},
		{
			scenario:      "異常系:ステータスコードが-100の場合インターナルサーバーエラーを生成",
			errMessage:    "エラーメッセージ",
			errStatusCode: -100,
		},
		{
			scenario:      "異常系:ステータスコードが450の場合インターナルサーバーエラーを生成\"",
			errMessage:    "エラーメッセージ",
			errStatusCode: 450,
		},
	}

	for _, c := range cases {
		t.Run(c.scenario, func(t *testing.T) {
			err := NewCustomError(c.errStatusCode, c.errMessage)
			if err == nil {
				t.Errorf("エラーが生成されていません。")
			}
			if err.StatusCode() == c.errStatusCode {
				t.Errorf("生成できないステータスコードが生成されています。 want:%d, receive:%d", c.errStatusCode, err.StatusCode())
			}
			if err.StatusCode() != http.StatusInternalServerError {
				t.Errorf("ステータスコードがエラー時のデフォルト値ではありません。 want:%d, receive:%d", http.StatusInternalServerError, err.StatusCode())
			}
			if err.Error() != "statusCode is not defined: "+c.errMessage {
				t.Errorf("エラーメッセージがエラー時のデフォルト値が含まれていません。 want:%s, receive:%s", "statusCode is not defined: "+c.errMessage, err.Error())
			}
		})
	}
}
