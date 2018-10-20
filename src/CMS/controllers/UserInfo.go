package controllers

import (
	"github.com/astaxie/beego"
	"CMS/models/DbModel"
	"time"
	"strconv"
	"github.com/astaxie/beego/orm"
	"strings"
)

type UserInfoController struct {
	beego.Controller
}

func (this *UserInfoController) Index() {
	this.TplName = "UserInfo/index.html"
}

//分页，搜索数据实体
type UserInfo struct {
	pageSize       int
	pageIndex      int
	SearchUserName string
	SearchRemark   string
	UserInfo       []DbModel.Userinfo
}

//添加数据
func (this *UserInfoController) AddUser() {
	o := orm.NewOrm()
	var userModel DbModel.Userinfo
	userModel.AddDate = time.Now()
	userModel.ModifDate = time.Now()
	userModel.DelFlag = 0 //正常状态
	userModel.UserName = this.GetString("UserName")
	userModel.UserPwd = this.GetString("UserPwd")
	userModel.Remark = this.GetString("Remark")
	_, err := o.Insert(&userModel)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"flag": "ok"}
	} else {
		this.Data["json"] = map[string]interface{}{"flag": "no"}
	}
	this.ServeJSON()
}

//查询数据
func (this *UserInfoController) GetUserList() {
	var u UserInfo
	u.pageSize, _ = strconv.Atoi(this.GetString("rows"))
	u.pageIndex, _ = strconv.Atoi(this.GetString("page"))
	u.SearchUserName = this.GetString("SearchUserName")
	u.SearchRemark = this.GetString("SearchRemark")
	beego.Info("拿到用户名了吗？", this.GetString("SearchUserName"))
	beego.Info("拿到备注了吗？", this.GetString("SearchRemark"))
	u.UserInfo = u.UserDatas()
	count := len(u.UserInfo)
	this.Data["json"] = map[string]interface{}{"rows": u.UserInfo, "total": count}
	this.ServeJSON()
}
func (this *UserInfo) UserDatas() []DbModel.Userinfo {
	o := orm.NewOrm()
	res := o.QueryTable("userinfo")

	if this.SearchUserName != "" {
		res = res.Filter("user_name__icontains", this.SearchUserName)
	}

	if this.SearchRemark != "" {
		res = res.Filter("remark__icontains", this.SearchRemark)
	}
	res = res.Filter("del_flag", 0)
	start := (this.pageIndex - 1) * this.pageSize

	res.OrderBy("id").Limit(this.pageSize, start).All(&this.UserInfo)
	return this.UserInfo
}

//删除用户
func (this *UserInfoController) RemoveUser() {
	ids := this.GetString("ids")
	o := orm.NewOrm()
	userIds := strings.Split(ids, ",")
	var userInfo DbModel.Userinfo
	//物理删除
	//for i := 0; i < len(userIds); i++ {
	//	userInfo.Id, _ = strconv.Atoi(userIds[i])
	//	_, err := o.Delete(&userInfo)
	//	if err != nil {
	//		beego.Info("删除错误", err)
	//		return
	//	}
	//}
	//软删除
	for i := 0; i < len(userIds); i++ {
		id, _ := strconv.Atoi(userIds[i])
		o.QueryTable("userinfo").Filter("id", id).One(&userInfo)
		userInfo.DelFlag = 1
		_, err := o.Update(&userInfo)
		if err != nil {
			beego.Info("更新失败", err)
			return
		}
	}
	this.Data["json"] = map[string]interface{}{"flag": "ok"}
	this.ServeJSON()
}

//修改数据
func (this *UserInfoController) EditUserInfo() {
	o := orm.NewOrm()
	var userInfo DbModel.Userinfo
	UserId := this.GetString("id")
	UserName := this.GetString("UserName")
	UserPwd := this.GetString("UserPwd")
	Remark := this.GetString("Remark")
	o.QueryTable("userinfo").Filter("id", UserId).One(&userInfo)
	userInfo.Remark = Remark
	userInfo.UserPwd = UserPwd
	userInfo.UserName = UserName
	userInfo.ModifDate = time.Now()
	_, err := o.Update(&userInfo)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"flag": "ok"}
	} else {
		this.Data["json"] = map[string]interface{}{"flag": "no"}
	}
	this.ServeJSON()
}

func (this *UserInfoController) ShowSetUserRole() {
	var user DbModel.Userinfo
	var userExtRoles []*DbModel.RoleInfo
	var allRolws []DbModel.RoleInfo
	o := orm.NewOrm()

	userId := this.GetString("userId")
	//查询用户
	o.QueryTable("userinfo").Filter("id", userId).One(&user)
	//关联用户已有角色
	o.LoadRelated(&user, "Roles")
	for _, role := range user.Roles {
		userExtRoles = append(userExtRoles, role)
	}
	//查询出所有角色
	o.QueryTable("role_info").Filter("del_flag", 0).All(&allRolws)
	this.Data["user"] = user
	this.Data["userExtRoles"] = userExtRoles
	this.Data["allRolws"] = allRolws

	//绑定视图
	this.TplName = "UserInfo/ShowSetUserRole.html"
}

