package main

import (
	_ "groupRegistration/routers"
	"github.com/astaxie/beego"
	"groupRegistration/models"
)

func main() {
	models.Init()
	beego.Run()
}

