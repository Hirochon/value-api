package initialization

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// TestDSNBuilderSuccess は、環境変数から受け取った値が正常かどうか確認
func TestDSNBuilderSuccess(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		dbms     string
		user     string
		password string
		host     string
		port     string
		database string
	}{
		{
			name:     "正常系: 環境変数が読み込めているかどうか",
			dbms:     "mysql",
			user:     os.Getenv("MYSQL_USER"),
			password: os.Getenv("MYSQL_PASSWORD"),
			host:     os.Getenv("MYSQL_HOST"),
			port:     os.Getenv("MYSQL_PORT"),
			database: os.Getenv("MYSQL_DATABASE"),
		},
		{
			name:     "正常系: 入れた値を構造体に格納できているかどうか",
			dbms:     "mysql",
			user:     "user",
			password: "password",
			host:     "host",
			port:     "port",
			database: "database",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := newDSNBuilder(c.dbms).
				setDSNUser(c.user).
				setDSNPassword(c.password).
				setDSNHost(c.host).
				setDSNPort(c.port).
				setDSNDatabase(c.database).
				checkError()
			if err != nil {
				t.Errorf("Error while building DSN: %v", err)
			}
		})
	}
}

// TestDSNBuilderFailed は、環境変数から受け取った値が異常だった場合に検知できるか確認
func TestDSNBuilderFailed(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		dbms     string
		user     string
		password string
		host     string
		port     string
		database string
	}{
		{
			name:     "異常系: dbmsが空文字列",
			dbms:     "",
			user:     "user",
			password: "password",
			host:     "host",
			port:     "port",
			database: "database",
		},
		{
			name:     "異常系: userが空文字列",
			dbms:     "mysql",
			user:     "",
			password: "password",
			host:     "host",
			port:     "port",
			database: "database",
		},
		{
			name:     "異常系: passwordが空文字列",
			dbms:     "mysql",
			user:     "user",
			password: "",
			host:     "host",
			port:     "port",
			database: "database",
		},
		{
			name:     "異常系: hostが空文字列",
			dbms:     "mysql",
			user:     "user",
			password: "password",
			host:     "",
			port:     "port",
			database: "database",
		},
		{
			name:     "異常系: portが空文字列",
			dbms:     "mysql",
			user:     "user",
			password: "password",
			host:     "host",
			port:     "",
			database: "database",
		},
		{
			name:     "異常系: databaseが空文字列",
			dbms:     "mysql",
			user:     "user",
			password: "password",
			host:     "host",
			port:     "port",
			database: "",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := newDSNBuilder(c.dbms).
				setDSNUser(c.user).
				setDSNPassword(c.password).
				setDSNHost(c.host).
				setDSNPort(c.port).
				setDSNDatabase(c.database).
				checkError()
			if err == nil {
				t.Error("空文字列があるにもかかわらずDSNが生成できている")
			}
		})
	}
}

// TestCombineDSNSuccess は、正常な値が入っている場合にDSNを構築できているか確認
func TestCombineDSNSuccess(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		dbms     string
		user     string
		password string
		host     string
		port     string
		database string
		expected string
	}{
		{
			name:     "環境変数ファイルでDSNを構築できるか",
			dbms:     "mysql",
			user:     os.Getenv("MYSQL_USER"),
			password: os.Getenv("MYSQL_PASSWORD"),
			host:     os.Getenv("MYSQL_HOST"),
			port:     os.Getenv("MYSQL_PORT"),
			database: os.Getenv("MYSQL_DATABASE"),
			expected: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE")),
		},
		{
			name:     "適当な値でDSNを構築できるか",
			dbms:     "mysql",
			user:     "user",
			password: "password",
			host:     "host",
			port:     "3306",
			database: "database",
			expected: "user:password@tcp(host:3306)/database?parseTime=true&loc=Asia%2FTokyo",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			valuesForMySQLDSN := &dsnEntity{
				user:     c.user,
				password: c.password,
				host:     c.host,
				port:     c.port,
				database: c.database,
				dbms:     c.dbms,
			}
			// たまにOSの環境変数が取れない場合があるので、少し待つ
			time.Sleep(time.Millisecond * 100)
			actual, err := valuesForMySQLDSN.combineMyENV()
			if err != nil {
				t.Errorf("Failed combine dsnEntity: %v", err)
			}
			if actual.toString() != c.expected {
				t.Errorf("dataSourceName is not correctly. Expected: %s, Actual: %s", c.expected, actual)
			}
		})
	}
}

// TestValidateMySQLDSNFailed は、異常な形式のDSNを検知できるか確認
func TestValidateMySQLDSNFailed(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name          string
		unexaminedDSN string
	}{
		{
			name:          "異常系: 誤っているDNSPart1",
			unexaminedDSN: "password@tcp(host:3306)/database?parseTime=true&loc=Asia%2FTokyo",
		},
		{
			name:          "異常系: 誤っているDNSPart2",
			unexaminedDSN: "user:@tcp(host:3306)/database?parseTime=true&loc=Asia%2FTokyo",
		},
		{
			name:          "異常系: 誤っているDNSPart3",
			unexaminedDSN: "@tcp(host:3306)/database?parseTime=true&loc=Asia%2FTokyo",
		},
		{
			name:          "異常系: 誤っているDNSPart4",
			unexaminedDSN: "user:password@tcp(host:3306)?parseTime=true&loc=Asia%2FTokyo",
		},
		{
			name:          "異常系: 誤っているDNSPart5",
			unexaminedDSN: "parseTime=true&loc=Asia%2FTokyo",
		},
		{
			name:          "異常系: 誤っているDNSPart6",
			unexaminedDSN: "",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := validateMySQLDSN(c.unexaminedDSN)
			if err == nil {
				t.Error("DSNの形式が異なっているにも関わらず、通ってしまう")
			}
		})
	}
}

// TestNewMySQLClient は、DSNによる接続の統合テスト！
func TestNewMySQLClientSuccess(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
	}{
		{
			name: "環境変数でMySQLへの接続を試みる",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := NewMySQLClient()
			if err != nil {
				t.Errorf("disconnect to MySQL by Environment: %v", err)
			}
		})
	}
}
func TestNewMySQLClientFailed(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		key   string
		value string
	}{
		{
			name:  "誤ったMYSQL_USERでMySQLへの接続が失敗する",
			key:   "MYSQL_USER",
			value: "",
		},
		{
			name:  "誤ったMYSQL_PORTでvalidateMySQLDSNが失敗する",
			key:   "MYSQL_PORT",
			value: "MALICE",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// 設定されている環境変数を一時保存
			saved := os.Getenv(c.key)
			// おかしな環境変数
			_ = os.Setenv(c.key, c.value)
			_, err := NewMySQLClient()
			if err == nil {
				t.Errorf("誤った環境変数で失敗しない事はないはず")
			}
			// 設定されている環境変数を一時保存解除
			_ = os.Setenv(c.key, saved)
		})
	}
}
