package controls

import (
	"github.com/astaxie/beego"
	"github.com/strive-after/go-kubernetes/module"
	"time"
	"github.com/strive-after/go-kubernetes/base/errors"
)
var (
	Secret string = "CMDB"
)
type LoginControllers struct {
	beego.Controller
}

type RegisterControllers struct {
	beego.Controller
}


//登陆显示页面
func (login *LoginControllers) LoginGet() {
	//获取cookie
	//email := login.Ctx.GetCookie("UserEmail")
	email, _:= login.Ctx.GetSecureCookie(Secret,"UserEmail")
	//true表示获取成功
	//if !ok {
	//	beego.Error("cookie获取失败")
	//	login.Redirect("/login",302)
	//}
	user := login.GetSession(email)

	//免密登陆  当cookie获取到值的时候  那么用cookie里面存的用户名称去获取session
	//session里面存了用户信息如果获取到说明用户是登陆状态  获取不到接口值为nil 那么就需要登陆
	if user == nil {
		login.TplName = `login.html`
	}else {
		login.Redirect("/",302)
	}
}

//登陆数据处理
func (login *LoginControllers) LoginPost()  {
	var (
		user module.User
		err error
	)
	//从前端获取用户输入信息
	if err := login.ParseForm(&user);err != nil {
		beego.Error("获取失败")
		login.Redirect("/LoginErr",302)
		return
	}
	if err = user.ComparePass(user.Password);err != nil{
		beego.Error("登陆失败")
		login.Redirect("/LoginErr",302)
		return
	}

	//如果记住用户名那么cookie保存时间为3600s
	err = user.Get("email",user.Email)
	if err != nil {
		beego.Error(err)
		login.Redirect("/LoginErr",302)
		return
	}

	login.Ctx.SetSecureCookie(Secret,"UserEmail",user.Email,time.Second*3600)
	//修改session
	login.SetSession(user.Email,user)
	login.Redirect("/",302)


}


//删除session
func (login *LoginControllers) LoginOut() {
	email := login.Ctx.GetCookie("UserEmail")
	login.DelSession(email)
	login.Ctx.SetSecureCookie(Secret,"UserEmail",email,-1)
	login.Redirect("/login",302)
}

//注册用户显示页面
func (reg *RegisterControllers) RegGet() {
	reg.TplName = `reg.html`
}
//用户注册数据处理
func (reg *RegisterControllers) RegPost() {
	var (
		inputuser module.User
		user  module.Operation   = new(module.User)
	)
	errs :=  errors.New()
	//将前端获取的数据直接赋值给user
	err := reg.ParseForm(&inputuser)
	if err != nil {
		errs.Add("Reg","注册失败")
	}
	err  = user.Add(&inputuser)
	if err != nil {
		errs.Add("Reg","注册失败")
	}

	if errs.HasErrors() {
		beego.Info(errs.Errors())
		reg.Data["err"] = errs
		reg.Redirect("/register",302)
	}

	reg.Redirect("/login",302)
}


func (login *LoginControllers)  Operation() {
	userEmail,ok  := login.Ctx.GetSecureCookie(Secret,"UserEmail")
	if !ok {
		beego.Error("获取cookie失败")
		return
	}
	user := login.GetSession(userEmail).(module.User)
	login.Data["UserName"] = user.Name
	login.TplName = `operation.html`
	login.Layout = `layout.html`
}


