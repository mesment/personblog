package db

import (
	"fmt"
	"log"
	"net/url"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mesment/personblog/pkg/setting"
)

var (
	DB *sqlx.DB
)

func Conn() *sqlx.DB  {
	return DB
}

func SetUp() {
	var err error
	dbType := setting.DBCfg.DBType
	user := setting.DBCfg.User
	passWd := setting.DBCfg.PassWord
	dbName := setting.DBCfg.DBName
	host := setting.DBCfg.Host
/*
	dns := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true",user,
		passWd,host,dbName)
	fmt.Printf("DNS:%s", dns)
	DB, err = sqlx.Open(dbType, dns)
*/

	uri :=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		user,
		passWd,
		host,
		dbName,
		url.QueryEscape("Asia/Shanghai" ),
	)
	fmt.Println("uri:",uri)
	DB, err = sqlx.Open(dbType, uri)
	if err != nil {
		log.Printf("连接数据库失败:%v",err)
		os.Exit(1)
	}

	err = DB.Ping()
	if err != nil {
		log.Printf("连接数据库失败:%v",err)
		os.Exit(1)
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return
}
