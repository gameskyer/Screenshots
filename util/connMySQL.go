package util
import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)
var DB *sqlx.DB
var err error
func init(){
	mysqlConnet := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local","root","123456","localhost:3306","screenshot")
	DB,err = sqlx.Open("mysql",mysqlConnet)
	// 最大连接数
	DB.SetMaxOpenConns(10)
	// 闲置连接数
	DB.SetMaxIdleConns(20)
	// 最大连接周期
	DB.SetConnMaxLifetime(100*time.Second)
	if err = DB.Ping(); nil != err {
		panic("数据库链接失败: " + err.Error())
	}
	//defer DB.Close()
}