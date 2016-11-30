package main

import (
	_ "appzone/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/logs"
	"fmt"
	"github.com/astaxie/beego/orm"
)
func init() {
	buildDBConfig()
	logs.SetLogger(logs.AdapterFile,`{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
}
func main() {
	beego.Run()
}
func buildDBConfig(){
	dbHost := beego.AppConfig.String("DBHost")
	dbPort := beego.AppConfig.String("DBPort")
	dbUser := beego.AppConfig.String("DBUser")
	dbPass :=beego.AppConfig.String("DBPass")
	dbName :=beego.AppConfig.String("DBName")
	dbDSN:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",dbUser,dbPass,dbHost,dbPort,dbName)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbDSN)
}
