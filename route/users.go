package route

import (
	"github.com/astaxie/beego"
	"github.com/strive-after/go-kubernetes/controls"
)

func init() {
	beego.AutoRouter(&controls.UserController{})
	//显示所有用户
	//beego.Router("/user/show",&controls.UserController{},"get:ShowUserGet")
	//查看用户信息
	//beego.Router("/user/info",&controls.UserController{},"get:UserInfo")
	//修改用户
	beego.Router("/user/change",&controls.UserController{},"get:ChangeUserGet;post:ChangeUserPost")
	//删除用户
	//beego.Router("/user/del",&controls.UserController{},"get:Del")
	//查看自己的信息
	beego.Router("/user/my/info",&controls.UserController{},"get:MyInfoGet;post:MyInfoPost")
	//修改当前用户密码
	beego.Router("/user/change/mypass",&controls.UserController{},"get:MyPassGet;post:MyPassPost")
	//管理员对用户做密码重置
	beego.Router("/user/change/userpass",&controls.UserController{},"get:UserPassGet;post:UserPassPost")
}