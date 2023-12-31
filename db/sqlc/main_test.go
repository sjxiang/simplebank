package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sjxiang/simplebank/util"
)


var testStore Store

// 主测试
func TestMain(m *testing.M) {
	// 跳转到，相对于当前 workspace 的根目录 /simplebank/db/sql -> /simplebank/db -> /simplebank 
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
