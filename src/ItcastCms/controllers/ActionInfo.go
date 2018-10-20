package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"strconv"
	"os"
	"ItcastCms/models"
	"github.com/astaxie/beego/orm"
)

type ActionInfoCtroller struct {
	beego.Controller
}

func (this *ActionInfoCtroller)Index()  {
	this.TplName="ActionInfo/Index.html"
}
//完成文件上传
func (this *ActionInfoCtroller)FileUp()  {
	f,h,err:=this.GetFile("fileUp")
	defer  f.Close()
	if err!=nil{
		this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件上传失败!!"}
	}else{
			//1:判断文件的类型
			  fileExt:=path.Ext(h.Filename)//文件名称
			  if fileExt==".jpg"||fileExt==".png"{
				  //2：判断文件的大小
				  if h.Size<50000000{
					  //3：创建不同的文件夹。
				 dir:="./static/fileUp/"+strconv.Itoa(time.Now().Year())+"/"+time.Now().Month().String()+"/"+strconv.Itoa(time.Now().Day())+"/"
				     //判断文件夹是否存在
				     _,err:=os.Stat(dir)
				     if err!=nil{//表示没有文件夹。
				     	os.MkdirAll(dir,os.ModePerm)
					 }
					  //4：文件的重命名。
					  newFileName:=strconv.Itoa(time.Now().Year())+time.Now().Month().String()+strconv.Itoa(time.Now().Day())+strconv.Itoa(time.Now().Hour())+strconv.Itoa(time.Now().Minute())+strconv.Itoa(time.Now().Nanosecond())
					  //构建一个完整上传路径。
					  fullDir:=dir+newFileName+fileExt
					  err1:=this.SaveToFile("fileUp",fullDir)//完成文件的保存。
					  if err1==nil{
						this.Data["json"]=map[string]interface{}{"flag":"ok","msg":fullDir}
					  }else {
						  this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件保存失败!!"}
					  }
				  }else{
					  this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件太大!!"}
				  }
			  }else{
				  this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件类型错误!!"}
			  }

		//5：实现文件的保存
		//6：返回消息。
		this.ServeJSON()


	}

}
//获取权限信息
func (this *ActionInfoCtroller)GetActionInfo()  {
  pageIndex,_:=this.GetInt("page")
  pageSize,_:=this.GetInt("rows")
  start:=(pageIndex-1)*pageSize
  o:=orm.NewOrm()
  var actions []models.ActionInfo
  o.QueryTable("action_info").Filter("del_flag",0).OrderBy("id").Limit(pageSize,start).All(&actions)
  count,_:=o.QueryTable("action_info").Filter("del_flag",0).Count()
  this.Data["json"]=map[string]interface{}{"rows":actions,"total":count}
  this.ServeJSON()
}


//完成权限的保存
func(this *ActionInfoCtroller) AddAction()  {
 var actionInfo=models.ActionInfo{}
 actionInfo.ActionTypeEnum,_=this.GetInt("ActionTypeEnum")
 actionInfo.MenuIcon=this.GetString("MenuIcon")
 actionInfo.Url=this.GetString("Url")
 actionInfo.ActionInfoName=this.GetString("ActionInfoName")
 actionInfo.IconWidth=0
 actionInfo.IconHeight=0
 actionInfo.HttpMethod=this.GetString("HttpMethod")
 actionInfo.Remark=this.GetString("Remark")
 actionInfo.DelFlag=0
 actionInfo.AddDate=time.Now()
 actionInfo.ModifDate=time.Now()
 o:=orm.NewOrm()
 _,err:=o.Insert(&actionInfo)
 if err==nil{
 	this.Data["json"]=map[string]interface{}{"flag":"ok"}
 }else{
	 this.Data["json"]=map[string]interface{}{"flag":"no"}
 }
 this.ServeJSON()


}
