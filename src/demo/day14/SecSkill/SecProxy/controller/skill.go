package controller

import "github.com/astaxie/beego"

type SkillController struct {
	//继承beego的controller一次使用beego的一些函数
	beego.Controller
}

func (skillC *SkillController) SecSkill() {
	//序列化
	skillC.Data["json"] = "secskill"
	//将数据转换成Json数据输出
	skillC.ServeJSON()
}
func (skillC *SkillController) SecInfo() {
	//序列化
	skillC.Data["json"] = "secInfo"
	//将数据转换成Json数据输出
	skillC.ServeJSON()
}
