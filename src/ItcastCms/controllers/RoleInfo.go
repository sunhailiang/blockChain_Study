package controllers

import (
	"github.com/astaxie/beego"
	"ItcastCms/models"
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	"strconv"
)

type RoleInfoController struct {
	beego.Controller
}

func (this *RoleInfoController)Index()  {
	this.TplName="RoleInfo/Index.html"
}
func (this *RoleInfoController)ShowAddRole()  {
	this.TplName="RoleInfo/ShowAddRole.html"
}
func (this *RoleInfoController)AddRole()  {
	var roleInfo=models.RoleInfo{}
	roleInfo.RoleName=this.GetString("roleName")
	roleInfo.Remark=this.GetString("roleRemark")
	roleInfo.DelFlag=0
	roleInfo.AddDate=time.Now()
	roleInfo.ModifDate=time.Now()
	o:=orm.NewOrm()
	_,err:=o.Insert(&roleInfo)
	if err==nil{
		this.Data["json"]=map[string]interface{}{"flag":"ok"}
	}else{
		this.Data["json"]=map[string]interface{}{"flag":"no"}
	}
	this.ServeJSON()

}
//获取角色信息
func (this *RoleInfoController)GetRoleInfo()  {
 pageIndex,_:=this.GetInt("page")
 pageSize,_:=this.GetInt("rows")
 start:=(pageIndex-1)*pageSize
 o:=orm.NewOrm()
 var roles []models.RoleInfo
 o.QueryTable("role_info").Filter("del_flag",0).OrderBy("Id").Limit(pageSize,start).All(&roles)
 count,_:=o.QueryTable("role_info").Filter("del_flag",0).Count()
 this.Data["json"]=map[string]interface{}{"rows":roles,"total":count}
 this.ServeJSON()
}
//展示角色已有的权限信息
func (this *RoleInfoController)ShowRoleAction()  {
  //1：接受传递过来的角色编号。
	roleId,_:=this.GetInt("roleId")
  //2:查询角色信息
  o:=orm.NewOrm()
  var roleInfo models.RoleInfo
  o.QueryTable("role_info").Filter("id",roleId).One(&roleInfo)
  //3:查询角色已经有的权限信息
  var roleExtActions []*models.ActionInfo //表示角色已经有的权限
  o.LoadRelated(&roleInfo,"Actions")
	for _,action:=range roleInfo.Actions{
		roleExtActions=append(roleExtActions,action)
	}
  //4：查询所有的权限信息。
  var allActions []models.ActionInfo
  o.QueryTable("action_info").Filter("del_flag",0).All(&allActions)
	this.Data["roleInfo"]=roleInfo
	this.Data["roleExtActions"]=roleExtActions
	this.Data["allActions"]=allActions
	this.TplName="RoleInfo/ShowSetRoleAction.html"
}
//完成角色对应的权限的分配
func (this *RoleInfoController)SetRoleAction()  {
  //1:接受角色id.
	roleId,_:=this.GetInt("roleId")
	//获取选中的权限的编号
	allKeys:=this.Ctx.Request.PostForm//获取所有的表单 map[string] []string
	var list[]int
	for key,_:=range allKeys {
		if strings.Contains(key,"cba_"){
			id:=strings.Replace(key,"cba_","",-1)
			strId,_:=strconv.Atoi(id)
			list=append(list,strId)
		}
	}
  //2:查询角色的信息
  o:=orm.NewOrm()
  var roleInfo models.RoleInfo
  o.QueryTable("role_info").Filter("id",roleId).One(&roleInfo)
  //3:根据查询出的角色信息，找出已有的权限信息看，并且全部的干掉。
  o.LoadRelated(&roleInfo,"Actions")
  m2m:=o.QueryM2M(&roleInfo,"Actions")
	for _,action:=range roleInfo.Actions {
		m2m.Remove(action)
	}
  //4：重新给角色分配权限。
    //根据list切片集合中存储的权限ID,查询出对应的权限信息，重新赋值给角色。
    var actionInfo models.ActionInfo
	for i:=0; i<len(list);i++  {
	o.QueryTable("action_info").Filter("id",list[i]).One(&actionInfo)
	m2m.Add(actionInfo)
	}
	this.Data["json"]=map[string]interface{}{"flag":"ok"}
	this.ServeJSON()

}