package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)
  var Db *sql.DB
//数据库的链接
func Connect()  {
	//项目配置
	config :=beego.AppConfig
	dbDriver :=config.String("db_driver")
	dbUser :=config.String("db_user")
	dbPassword :=config.String("db_password")
	dbIp:=config.String("db_ip")
	dbName :=config.String("db_name")
	fmt.Println(dbDriver,dbUser,dbPassword,dbIp,dbName)
	//连接数据库
	connUrl :=dbUser+":"+dbPassword +"@tcp("+dbIp+")/"+dbName+"?"
	db,err :=sql.Open(dbDriver,connUrl)
	if err!=nil{
		panic("数据库链接错误，请检查配置")
	}
	Db= db
}