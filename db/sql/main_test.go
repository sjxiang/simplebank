package db

import (
	"log"
	"os"
	"testing"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)


var testData *Data

// 主测试
func TestMain(m *testing.M) {
	
	// driver + dsn
	connector, err := mysql.NewConnector(&mysql.Config{
		User:      "root",
		Passwd:    "secret",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "simple_bank",
		Collation: "utf8mb4_general_ci",
		ParseTime: true,
	})
	if err != nil {
		panic(err)
	}
	
	db := sql.OpenDB(connector)
	if err := db.Ping(); err != nil {
		log.Fatal("failed to connect database:", err)
	}

	testData = NewData(db)
	os.Exit(m.Run())
}