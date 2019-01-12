package main

import (
	_ "groupRegistration/routers"
	"github.com/astaxie/beego"
	"groupRegistration/models"
	"groupRegistration/lib"
)

func main() {
	models.Init()
	lib.InitLogger()
	beego.Run()
}

