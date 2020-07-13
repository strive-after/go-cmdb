package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"

	_ "github.com/strive-after/go-cmdb/module"
	_ "github.com/strive-after/go-cmdb/route"
)




func main() {
	//未注册  必须注册才能用 有时候session获取的时候  它会提示  gob: name not registered for interface: "github.com/strive-after/go-kubernetes/module.User"
	//那么就注册一下
	beego.Run()
}
