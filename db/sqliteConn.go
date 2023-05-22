package db

import (
	"fmt"
	logs "github.com/cihub/seelog"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var DB *gorm.DB

const (
	dbDriverName    = "sqlite3"
	dbName          = "./db/data.db3"
	SCRENSHOTSTABLE = "screenshot"
)

func init() {
	// github.com/mattn/go-sqlite3
	if !Exists("./db") {
		os.Mkdir("./db", 0755)
	}
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		PrepareStmt: true, Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		logs.Error("Open mysql failed,err:%v\n", err)
		panic(err)
	}
	DB = db
	sqlDB, err := DB.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(100)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(1000)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute)
}

func CheckErr(e error) bool {
	if e != nil {
		logs.Error(e)
		return true
	}
	return false
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) { //判断为true,说明文件或文件夹存在
			return true
		}
		return false
	}
	return true
}
