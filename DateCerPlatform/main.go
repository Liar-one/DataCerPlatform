package main

import (
	"DateCerPlatform/db_mysql"
	_ "DateCerPlatform/routers"
	"github.com/astaxie/beego"
)

func main() {
	//接收数据库
      db_mysql.Connect()
    //设置静态资源文件路径设置
    beego.SetStaticPath("/js","/static/js")
    beego.SetStaticPath("/css","/static/css")
    beego.SetStaticPath("/img","/static/img")
	beego.Run()
}

