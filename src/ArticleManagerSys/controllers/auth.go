package controllers

import (
	"github.com/astaxie/beego"
	"ArticleManagerSys/models"
)

type AuthController struct {
	beego.Controller
}

//登陆
func (this *AuthController) LoginView() {
	this.TplName = "login.html"
}
func (this *AuthController) Login() {
	var userObj models.UserModel
	userObj.Name = this.GetString("userName")
	userObj.Pwd = this.GetString("passWord")
	var res = userObj.Login(userObj)
	if res == 100 {
		this.SetSession("user", userObj.Name)
		this.Redirect("/", 302)
	} else {
		this.Data["loginErr"] = "账号或者密码错误"
		this.TplName = "login.html"
	}
}

//注册
func (this *AuthController) RegistView() {
	this.TplName = "register.html"
}

func (this *AuthController) Regist() {
	var userObj models.UserModel
	userObj.Name = this.GetString("userName")
	userObj.Pwd = this.GetString("password")
	var res = userObj.Regist(userObj)
	if res == 000 {
		this.Redirect("/", 302)
	} else if res == 001 {
		this.Data["errInfo"] = "后台出错：code:001"
		this.TplName = "register.html"
	} else {
		this.Data["errInfo"] = "账号已经存在"
		this.TplName = "register.html"
	}
}

//退出
func (this *AuthController) Logout() {
	beego.Info("get session ", this.GetSession("user"))
	this.DelSession("user")
	this.Data["json"]="{"+"code"+":"+"302"+"}"
	this.ServeJSON()
}
