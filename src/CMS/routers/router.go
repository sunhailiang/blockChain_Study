package routers

import (
	"CMS/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/Admin/UserInfo/Index", &controllers.UserInfoController{},"get:Index")
	beego.Router("/UserInfo/AddUser", &controllers.UserInfoController{},"post:AddUser")
	beego.Router("/Admin/UserInfo/GetUserList", &controllers.UserInfoController{},"post:GetUserList")
	beego.Router("/Admin/UserInfo/RemoveUser", &controllers.UserInfoController{},"post:RemoveUser")
	beego.Router("/Admin/UserInfo/EditUserInfo", &controllers.UserInfoController{},"post:EditUserInfo")

//=========================角色分配
	beego.Router("/Admin/UserInfo/ShowSetUserRole",&controllers.UserInfoController{},"get:ShowSetUserRole")
	beego.Router("/Admin/UserInfo/ShowSetUserRole",&controllers.UserInfoController{},"get:ShowSetUserRole")

}
