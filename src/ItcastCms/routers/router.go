package routers

import (
	"ItcastCms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/Admin/UserInfo/Index",&controllers.UserInfoController{},"get:Index")
    beego.Router("/Admin/UserInfo/AddUser",&controllers.UserInfoController{},"post:AddUser")
	beego.Router("/Admin/UserInfo/GetUserInfo",&controllers.UserInfoController{},"post:GetUserInfo")
beego.Router("/Admin/UserInfo/DeleteUser",&controllers.UserInfoController{},"post:DeleteUser")
beego.Router("/Admin/UserInfo/ShowEditUser",&controllers.UserInfoController{},"post:ShowEditUser")
	beego.Router("/Admin/UserInfo/ShowSetUserRole",&controllers.UserInfoController{},"get:ShowSetUserRole")
beego.Router("/Admin/UserInfo/SetUserRole",&controllers.UserInfoController{},"post:SetUserRole")
    //-------------------------------角色管理--------------------------------------->
beego.Router("/Admin/RoleInfo/Index",&controllers.RoleInfoController{},"get:Index")
beego.Router("/Admin/RoleInfo/ShowAddRole",&controllers.RoleInfoController{},"get:ShowAddRole")
beego.Router("/Admin/RoleInfo/AddRole",&controllers.RoleInfoController{},"post:AddRole")
beego.Router("/Admin/RoleInfo/GetRoleInfo",&controllers.RoleInfoController{},"post:GetRoleInfo")
 beego.Router("/Admin/RoleInfo/ShowRoleAction",&controllers.RoleInfoController{},"get:ShowRoleAction")
beego.Router("/Admin/RoleInfo/SetRoleAction",&controllers.RoleInfoController{},"post:SetRoleAction")
//------------------------------------权限管理---------------------------------------------------------
beego.Router("/Admin/ActionInfo/Index",&controllers.ActionInfoCtroller{},"get:Index")
beego.Router("/Admin/ActionInfo/FileUp",&controllers.ActionInfoCtroller{},"post:FileUp")
beego.Router("/Admin/ActionInfo/AddAction",&controllers.ActionInfoCtroller{},"post:AddAction")
beego.Router("/Admin/ActionInfo/GetActionInfo",&controllers.ActionInfoCtroller{},"post:GetActionInfo")
}
