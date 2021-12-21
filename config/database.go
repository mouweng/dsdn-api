package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

const DbName = "zoneDB"

var db map[string]*sql.DB = make(map[string]*sql.DB)
var dbMutex sync.RWMutex
var (
	dbUsername = map[string]string{}
	dbPassword = map[string]string{}
	dbAddress  = map[string]string{}
	dbName     = map[string]string{}
)

// initializeDatabase 初始化数据库连接信息
func initializeDatabase(config Config) {
	switch config.String("mode") {
	case "test":
		dbAddress[DbName] = "9.134.190.21:3306"
		dbUsername[DbName] = "root"
		dbPassword[DbName] = "123456"
		dbName[DbName] = "test"
	case "release":
		dbAddress[DbName] = "9.134.190.21:3306"
		dbUsername[DbName] = "root"
		dbPassword[DbName] = "123456"
		dbName[DbName] = "test"
	default:
		dbAddress[DbName] = "127.0.0.1:3306"
		dbUsername[DbName] = "root"
		dbPassword[DbName] = "123456"
		dbName[DbName] = "test"
	}
}

func GetDBConnect(dbname string) func() *sql.DB {
	return func() *sql.DB {
		dbMutex.RLock()
		conn := db[dbname]
		if conn != nil {
			dbMutex.RUnlock()
			return conn
		}
		dbMutex.RUnlock()

		dbMutex.Lock()
		defer dbMutex.Unlock()
		conn = db[dbname]
		if conn != nil {
			return conn
		}

		connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8&timeout=5s&readTimeout=30s&writeTimeout=30s&sql_mode=%%27ANSI%%27", dbUsername[dbname], dbPassword[dbname], dbAddress[dbname], dbName[dbname])
		var err error
		conn, err = sql.Open("mysql", connStr)
		if err != nil {
			log.Fatal("Connect to database failed: ", connStr)
			log.Panicln(err)
		}
		conn.SetMaxIdleConns(10)
		db[dbname] = conn
		return conn
	}
}
