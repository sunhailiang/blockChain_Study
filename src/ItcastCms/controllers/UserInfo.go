package controllers

import "github.com/astaxie/beego"
import (
	"ItcastCms/models"
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)
type UserInfoController struct {
	beego.Controller
}
//存储的是搜索的条件
type UserData struct {
	PageIndex int
	PageSize int
	Name string
	Remark string
	TotalCount int64
}

func(this *UserInfoController) Index()  {
	this.TplName="UserInfo/Index.html"
}
//获取数据表中的数据，返回前端表格。
func(this *UserInfoController) GetUserInfo()  {
  //1:接收前端发送过来的当前页码，与每页显示的记录数。
   pageIndex,_:=strconv.Atoi(this.GetString("page"))//当前页码。
   pageSize,_:=strconv.Atoi(this.GetString("rows"))//每页显示记录数。
   //获取搜索的用户名和备注
   name:=this.GetString("name")
   remark:=this.GetString("remark")

   //实现搜索
   //创建一个搜索的对象（UserData）,然后将搜索条件和分页的条件数据，赋值给搜索对象中的属性。
   //将该对象作为SearchUserData的参数。
   var userSearchData=UserData{}
   userSearchData.Remark=remark
   userSearchData.PageSize=pageSize
   userSearchData.PageIndex=pageIndex
   userSearchData.Name=name
	serverData:=userSearchData.SearchUserData()
	this.Data["json"]=map[string]interface{}{"rows":serverData,"total":userSearchData.TotalCount}
	this.ServeJSON()
  //2:实现分页查询数据
  //Limit(第一个参数，第二个参数)
  //第一个参数：获取多少条数据。
  //第二个参数：从哪儿开始取。
  /*
  start:=(pageIndex-1)*pageSize
  var users []models.UserInfo
  o:=orm.NewOrm()
  o.QueryTable("user_info").Filter("del_flag",0).OrderBy("id").Limit(pageSize,start).All(&users)
  count,_:=o.QueryTable("user_info").Filter("del_flag",0).Count()//获取总的记录数。
	  //3:将数据构成成json，返回前端表格。
	  //注意：数据的KEY必须是rows,总的记录必须是total.
	  this.Data["json"]=map[string]interface{}{"rows":users,"total":count}
	  this.ServeJSON()*/
}
func(this *UserData) SearchUserData()[]models.UserInfo  {
  //构建搜索条件。
  o:=orm.NewOrm()
  temp:=o.QueryTable("user_info")
	if this.Name!="" {		//"a"

		temp=temp.Filter("user_name__icontains",this.Name)
	}
	if this.Remark!=""{
		temp=temp.Filter("remark__icontains",this.Remark)
	}
	temp=temp.Filter("del_flag",0)
	//实现分页了
	this.TotalCount,_=temp.Count()//总的记录数。
	start:=(this.PageIndex-1)*this.PageSize
	var users[]models.UserInfo
	temp.OrderBy("Id").Limit(this.PageSize,start).All(&users)
	return users


}

//完成用户数据的添加
func(this* UserInfoController) AddUser()  {
	var userInfo=models.UserInfo{}
	userInfo.UserName=this.GetString("userName")//接收用户名
	userInfo.UserPwd=this.GetString("userPwd")
	userInfo.Remark=this.GetString("userRemark")
	userInfo.ModifDate=time.Now()
	userInfo.AddDate=time.Now()
	userInfo.DelFlag=0//表示正常，1表示表示软删除。
	o:=orm.NewOrm()
	_,err:=o.Insert(&userInfo)
	if err==nil{
		//Data中的key必须为"json"
		this.Data["json"]=map[string]interface{}{"flag":"ok"}

	}else{
		this.Data["json"]=map[string]interface{}{"flag":"no"}
	}
	//怎样将数据生成JSON.
	this.ServeJSON()

}
//删除记录
func (this *UserInfoController)DeleteUser()  {
	//1:接收传递的数据的Id值。
	ids:=this.GetString("strId")
	//2:按照逗号进行分隔。
	strIds:=strings.Split(ids,",")
	//3:删除指定的数据。
	  //3.1: 将切片中存储的Id有字符串转成整型。
	  o:=orm.NewOrm()
	var userInfo=models.UserInfo{}
	for i:=0; i<len(strIds); i++ {
		id,_:=strconv.Atoi(strIds[i])
		//3.2删除数据
		userInfo.Id=id
		o.Delete(&userInfo)
	}
	this.Data["json"]=map[string]interface{}{"flag":"ok"}
	this.ServeJSON()

}

//展示要编辑的数据
func (this *UserInfoController)ShowEditUser()  {
	userId,_:=this.GetInt("userId")
	var userInfo models.UserInfo
	o:=orm.NewOrm()
	o.QueryTable("user_info").Filter("id",userId).One(&userInfo)
	this.Data["json"]=map[string]interface{}{"userInfo":userInfo}
	this.ServeJSON()
}
//完成数据的更新
func (this *UserInfoController)EditUser()  {
	//1接收到要更新的数据
	var userInfo=models.UserInfo{}
	userInfo.UserName=this.GetString("userEditName")
	userInfo.UserPwd=this.GetString("userEditPwd")
	userInfo.Remark=this.GetString("userEditRemark")
	userInfo.Id,_=this.GetInt("userEidtId")
   //2:

}

//展示要给用户分配的角色信息
func (this *UserInfoController)ShowSetUserRole()  {
//1：接收传递过来的用户编号
  userId,_:=this.GetInt("userId")
//2:查询出用户已经有的角色。
var userInfo models.UserInfo
o:=orm.NewOrm()
o.QueryTable("user_info").Filter("id",userId).One(&userInfo)
var userExtRoles []*models.RoleInfo//表示用户已有的角色。
o.LoadRelated(&userInfo,"Roles")//在这直接根据建立的关系（多对多）查询数据，将数据给了Roles
	for _,role:=range userInfo.Roles{
		userExtRoles=append(userExtRoles,role)
	}
//3:查询出所有角色。
 var allRoles[]models.RoleInfo
 o.QueryTable("role_info").Filter("del_flag",0).All(&allRoles)
 this.Data["allRoles"]=allRoles
 this.Data["userExtRoles"]=userExtRoles
 this.Data["userInfo"]=userInfo

 this.TplName="UserInfo/ShowSetUserRole.html"
}
//完成用户角色的分配
func(this *UserInfoController) SetUserRole()  {
	allKeys:=this.Ctx.Request.PostForm//获取表单中的数据，返回的map key为name的值。value,是用户在表单中输入的值。
	//获取了用户选择的所有角色的编号。
	//需要将name前面的前缀替换。
	var list[]int
	for key,_:= range allKeys{
		if strings.Contains(key,"cba_"){
			id:=strings.Replace(key,"cba_","",-1)
			roleId,_:=strconv.Atoi(id)
			list=append(list,roleId)
		}
	}
	//给用户分配角色，一定要知道用户的编号。
	userId,_:=this.GetInt("userId")
  //怎么给用户完成角色的分配？
  //获取用户已有的角色，并且全部删除掉，然后再重新分配。
   var userInfo models.UserInfo
     o:=orm.NewOrm()
     o.QueryTable("user_info").Filter("id",userId).One(&userInfo)
   o.LoadRelated(&userInfo,"Roles");
   m2m:=o.QueryM2M(&userInfo,"Roles")//指定userInfo 与Roles的关系。该对象用来操作多对多的关系。
  //删除用户已经有的角色.
	for _,role:= range userInfo.Roles{
		m2m.Remove(role)//删除的是中间表中的数据。
	}
	//重新插入。
	//从list集合中获取了要给用户分配的角色的信息。
	var roleInfo models.RoleInfo
	for i:=0;i<len(list) ;i++  {
		o.QueryTable("role_info").Filter("id",list[i]).One(&roleInfo)
		m2m.Add(roleInfo)

	}
	this.Data["json"]=map[string]interface{}{"flag":"ok"}
	this.ServeJSON()

}