package models

import "github.com/astaxie/beego/orm"
import _"github.com/go-sql-driver/mysql"

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.233.144:3306)/beego?charset=utf8&loc=Local")
	//映射对象
	orm.RegisterModel(new(UserModel),new(ArticleModel),new(ArticleTypeModel))
	//创建表
	orm.RunSyncdb("default", false, true)
}
