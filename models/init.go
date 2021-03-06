package models

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
)
/**
数据库链接
*/
func Init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	//timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	//拼接数据库链接
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Local"
	fmt.Println(dsn)
	//if timezone == "" {
	//	dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	//}
	maxIdle := 30
	maxConn := 300
	orm.RegisterDataBase("default","mysql",dsn,maxIdle,maxConn)
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModel(new(PinActivity),new(PinEvent))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return name
}
