package controls

import (
	"github.com/astaxie/beego"
	"github.com/strive-after/go-kubernetes/base/errors"
	"github.com/strive-after/go-kubernetes/module"
	"time"
)
var (
	Secret string = "CMDB"
)
type AuthController struct {
	beego.Controller
}

//type RegisterController struct {
//	beego.Controller
//}

type Operation struct {
	beego.Controller
}

////登陆显示页面
//func (login *AuthController) Login() {
//	//获取cookie
//	email, _:= login.Ctx.GetSecureCookie(Secret,"UserEmail")
//	user := login.GetSession(email)
//	//免密登陆  当cookie获取到值的时候  那么用cookie里面存的用户名称去获取session
//	//session里面存了用户信息如果获取到说明用户是登陆状态  获取不到接口值为nil 那么就需要登陆
//	if user == nil {
//		login.TplName = `login.html`
//	}else {
//		login.Redirect("/",302)
//	}
//}

//登陆数据处理
func (login *AuthController) Login()  {
	var (
		user module.User
		err error
	)
	errs := errors.New()
	//如果是get直接加载页面
	//如果是post 做数据 处理
	if login.Ctx.Input.IsPost() {
		if err := login.ParseForm(&user);err != nil {
			errs.Add("Login","登陆失败")
			return
		}
		if err = user.ComparePass(user.Password);err != nil{
			errs.Add("Login","登陆失败")
			beego.Error("登陆失败")
			login.Redirect("/auth/login",302)
			return
		}

		//如果记住用户名那么cookie保存时间为3600s
		err = user.Get("email",user.Email)
		if err != nil {
			errs.Add("Login","登陆失败")
			beego.Error(err)
			login.Redirect("/auth/login",302)
			return
		}
		login.Ctx.SetSecureCookie(Secret,"UserEmail",user.Email,time.Second*3600)
		login.SetSession(user.Email,user)
		login.Redirect("/",302)
		return
	}
	login.TplName = `login.html`
}


//删除session

func (login *AuthController) Out() {
	email := login.Ctx.GetCookie("UserEmail")
	login.DelSession(email)
	login.Ctx.SetSecureCookie(Secret,"UserEmail",email,-1)
	login.Redirect("/auth/login",302)
}

//注册用户显示页面
//func (reg *RegisterController) RegGet() {
//	reg.TplName = `reg.html`
//}
//用户注册数据处理
func (reg *AuthController) Reg() {
	var (
		inputuser module.User
		user  module.Operation   = new(module.User)
	)
	errs :=  errors.New()
	if reg.Ctx.Input.IsPost() {
		//将前端获取的数据直接赋值给user
		err := reg.ParseForm(&inputuser)
		if err != nil {
			beego.Error(err)
			errs.Add("Reg","注册失败")
		}
		err  = user.Add(&inputuser)
		if err != nil {
			beego.Error(err)
			errs.Add("Reg","注册失败")
		}
		if errs.HasErrors() {
			beego.Info(errs.Errors())
			reg.Data["err"] = errs.Errors()
			reg.Redirect("/register",302)
			return
		}
		reg.Redirect("/auth/login",302)
		return
	}
	reg.TplName = `reg.html`
}


func (login *Operation)  Get() {
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


