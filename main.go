package main

import (
	"fmt"
	"kindergarten-service-go/controllers"
	_ "kindergarten-service-go/routers"
	"kindergarten-service-go/services"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var host = beego.AppConfig.String("DB_HOST")
var port = beego.AppConfig.String("DB_PORT")
var username = beego.AppConfig.String("DB_USERNAME")
var password = beego.AppConfig.String("DB_PASSWORD")
var database = beego.AppConfig.String("DB_DATABASE")
var connection = beego.AppConfig.String("DB_CONNECTION")

func init() {
	orm.RegisterDataBase("default", connection, username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&loc=Local")
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}

	filename := time.Now().Format("20060102")
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"storage/logs/%s.log"}`, filename))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	kindergartenServer := services.KindergartenServer{}
	kindergartenServer.Init()
	beego.Run()

}
