/*
mysql_client.go は、MySQL クライアントを構築するためのパッケージ
*/

package initialization

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"regexp"
	"time"
)

const (
	maxIdleConns    = 10
	maxOpenConns    = 10
	connMaxLifetime = 60 * time.Second
)

type dsnEntity struct {
	user     string
	password string
	host     string
	port     string
	database string
	dbms     string
}

type dsnBuildEntity struct {
	user     string
	password string
	host     string
	port     string
	database string
	dbms     string
	err      error
}

func newDSNBuilder(dbms string) *dsnBuildEntity {
	if dbms == "" {
		return &dsnBuildEntity{err: fmt.Errorf("dbms is empty")}
	}
	return &dsnBuildEntity{
		dbms: dbms,
	}
}

func (dsn *dsnBuildEntity) setDSNUser(s string) *dsnBuildEntity {
	if dsn.err != nil {
		return dsn
	}
	if dsn.user = s; dsn.user == "" {
		dsn.err = fmt.Errorf("%s: dsn.user is empty", dsn.dbms)
	}
	return dsn
}

func (dsn *dsnBuildEntity) setDSNPassword(s string) *dsnBuildEntity {
	if dsn.err != nil {
		return dsn
	}
	if dsn.password = s; dsn.password == "" {
		dsn.err = fmt.Errorf("%s: dsn.password is empty", dsn.dbms)
	}
	return dsn
}

func (dsn *dsnBuildEntity) setDSNHost(s string) *dsnBuildEntity {
	if dsn.err != nil {
		return dsn
	}
	if dsn.host = s; dsn.host == "" {
		dsn.err = fmt.Errorf("%s: dsn.host is empty", dsn.dbms)
	}
	return dsn
}

func (dsn *dsnBuildEntity) setDSNPort(s string) *dsnBuildEntity {
	if dsn.err != nil {
		return dsn
	}
	if dsn.port = s; dsn.port == "" {
		dsn.err = fmt.Errorf("%s: dsn.port is empty", dsn.dbms)
	}
	return dsn
}

func (dsn *dsnBuildEntity) setDSNDatabase(s string) *dsnBuildEntity {
	if dsn.err != nil {
		return dsn
	}
	if dsn.database = s; dsn.database == "" {
		dsn.err = fmt.Errorf("%s: dsn.database is empty", dsn.dbms)
	}
	return dsn
}

func (dsn *dsnBuildEntity) checkError() (*dsnEntity, error) {
	if dsn.err != nil {
		return nil, dsn.err
	}
	return &dsnEntity{
		user:     dsn.user,
		password: dsn.password,
		host:     dsn.host,
		port:     dsn.port,
		database: dsn.database,
		dbms:     dsn.dbms,
	}, nil
}

type dataSourceName string

func (dataSourceName dataSourceName) toString() string {
	return string(dataSourceName)
}

// validateMySQLDSN は、正規表現によって正常なDSNかどうかを判断する
func validateMySQLDSN(combinedStringForMySQLDSN string) (dataSourceName, error) {
	r, err := regexp.Compile(`.+:.+@tcp\(.+:\d+\)/.+`)
	if err != nil {
		return "", err
	}
	if !r.MatchString(combinedStringForMySQLDSN) {
		return "", fmt.Errorf("%s is not a valid MySQL DSN", combinedStringForMySQLDSN)
	}
	return dataSourceName(combinedStringForMySQLDSN), nil
}

func (dsn dsnEntity) combineMyENV() (dataSourceName, error) {
	combinedMySQLDSN := dsn.user + ":" + dsn.password + "@tcp(" + dsn.host + ":" + dsn.port + ")/" + dsn.database + "?parseTime=true&loc=Asia%2FTokyo"
	return validateMySQLDSN(combinedMySQLDSN)
}

// NewMySQLClient は、MySQL クライアントを構築
func NewMySQLClient() (*sqlx.DB, error) {
	valuesForMySQLDSN, err := newDSNBuilder("mysql").
		setDSNUser(os.Getenv("MYSQL_USER")).
		setDSNPassword(os.Getenv("MYSQL_PASSWORD")).
		setDSNHost(os.Getenv("MYSQL_HOST")).
		setDSNPort(os.Getenv("MYSQL_PORT")).
		setDSNDatabase(os.Getenv("MYSQL_DATABASE")).
		checkError()
	if err != nil {
		return nil, err
	}

	dataSourceName, err := valuesForMySQLDSN.combineMyENV()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Open("mysql", dataSourceName.toString())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	return db, nil
}
