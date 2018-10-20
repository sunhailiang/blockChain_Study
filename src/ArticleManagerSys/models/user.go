package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type UserModel struct {
	Id         int             `orm:"pk;auto"`
	Name       string          `orm:"size(20)"`
	Pwd        string          `orm:"size(30)"`
	CreateTime time.Time       `orm:"auto_now"`
	Articles   []*ArticleModel `orm:"rel(m2m)"`
}

//注册
func (this *UserModel) Regist(userDate UserModel) (res int) {
	o := orm.NewOrm()
	this.Name = userDate.Name
	this.Pwd = userDate.Pwd
	if o.Read(this, "name") == nil {
		return 100
	} else {
		_, err := o.Insert(this)
		if err != nil {
			beego.Info("数据插入失败", err)
			return 001
		} else {
			return 000
		}
	}
}

//登陆
func (this *UserModel) Login(userDate UserModel) (res int) {
	o := orm.NewOrm()
	this.Name = userDate.Name
	this.Pwd = userDate.Pwd
	if o.Read(this, "name", "pwd") == nil {
		return 100
	} else {
		return 101
	}

}
