package util

import (
	"Screenshots/module"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	DB        *sqlx.DB
	err       error
	sqlConfig module.SqlConfig
)

func init() {
	sqlConfig = Config.SqlConfig
	mysqlConnet := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sqlConfig.DbUser, sqlConfig.DbPassword, sqlConfig.DbAddress, sqlConfig.DbName)
	DB, err = sqlx.Open("mysql", mysqlConnet)
	// 最大连接数
	DB.SetMaxOpenConns(10)
	// 闲置连接数
	DB.SetMaxIdleConns(20)
	// 最大连接周期
	DB.SetConnMaxLifetime(100 * time.Second)
	if err = DB.Ping(); nil != err {
		panic("数据库链接失败: " + err.Error())
	}
	//defer DB.Close()
}
