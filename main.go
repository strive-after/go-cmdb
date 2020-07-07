package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/strive-after/go-kubernetes/module"
	_ "github.com/strive-after/go-kubernetes/module"
	_ "github.com/strive-after/go-kubernetes/route"
	"strconv"
	"time"
)

func StopTimeFormat(stoptime time.Time) string{
	stop := stoptime.Format("2006-01-02")

	if stop == "0001-01-01" {
		return "无"
	}
	return stop
}

func IndexLeft(index int) string{
	num := strconv.Itoa(index-1)
	return  num
}
func IndexRight(index int) string{
	num := strconv.Itoa(index+1)
	return  num
}

func Role(num int) string{
	switch num {
	case 0:
		return "普通用户"
	case 1:
		return "管理员"
	case 2:
		return "超级管理员"
	}
	return ""
}

func main() {
	//未注册  必须注册才能用 有时候session获取的时候  它会提示  gob: name not registered for interface: "github.com/strive-after/go-kubernetes/module.User"
	//那么就注册一下
	gob.Register(module.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("redisurl")
	beego.AddFuncMap("TimeForMat",StopTimeFormat)
	beego.AddFuncMap("Roles",module.Role)
	beego.AddFuncMap("Left",IndexLeft)
	beego.AddFuncMap("Right",IndexRight)
	err := logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)  //separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	if err != nil {
		beego.Error(err)
		return
	}
	beego.Run()
}
