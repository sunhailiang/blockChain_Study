package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"math"
)

type ArticleModel struct {
	Id          int               `orm:"pk;auto"`
	ArticleName string            `orm:"size(30)"`
	CreateTime  time.Time         `orm:"auto_now"`
	Count       int               `orm"default(0);null"`
	Content     string            `orm:size(1000)`
	Img         string            `orm:size(5000)`
	Type        string            `orm:size(20)`
	ArticleType *ArticleTypeModel `orm:"rel(fk)"`
	UserModel   []*UserModel      `orm:"reverse(many)"`
}

type ArticleTypeModel struct {
	Id       int
	TypeName string          `orm:"size(20)"`
	Articles []*ArticleModel `orm:"reverse(many)"`
}

func (this *ArticleModel) AddRrticle(obj ArticleModel) (res int) {
	o := orm.NewOrm()
	this.ArticleName = obj.ArticleName
	this.Content = obj.Content
	this.Count = obj.Count
	this.Img = obj.Img
	this.Type = obj.Type
	//检索type
	var typeObj ArticleTypeModel
	typeObj.TypeName = this.Type
	beego.Info("有没有类型", this.Type)
	err := o.Read(&typeObj, "TypeName")
	if err != nil {
		beego.Info("获取文章类型失败", err)
		return
	}
	//添加关联数据
	this.ArticleType = &typeObj

	_, err = o.Insert(this)
	if err != nil {
		beego.Info("文章插入失败", err)
		return 501
	}
	return 200
}

//分页
type Page struct {
	PageNum    int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}
type ResPageObj struct {
	Article []ArticleModel
	Page    Page
	Type    []ArticleTypeModel
}

func (this *ArticleModel) GetArticleList(index int, selectType string) ResPageObj {
	//所有返回对象的集合
	var obj ResPageObj
	//查询文章类型
	var typeObj ArticleTypeModel
	var resTypesObj []ArticleTypeModel
	resTypesObj = typeObj.GetArticleTypeList()
	obj.Type = resTypesObj
	//查询数据
	o := orm.NewOrm()

	var article []ArticleModel
	res := o.QueryTable("ArticleModel")
	count, _ := res.RelatedSel("ArticleType").Count()
	beego.Info("有数据吗？", count)
	pageSize := 5
	pageIndex := index
	start := pageSize * (pageIndex - 1)
	//分页参数
	obj.Page.TotalPage = int(math.Ceil(float64(count) / float64(pageSize)))
	obj.Page.TotalCount = int(count)
	obj.Page.PageNum = pageIndex
	res.Limit(pageSize, start).RelatedSel("ArticleType").All(&article)
	if pageIndex < 2 {
		obj.Page.FirstPage = true
	}
	if pageIndex >= obj.Page.TotalPage {
		obj.Page.LastPage = true
	}
	if selectType == "" {
		beego.Info("下拉框传递数据失败")
		res.Limit(pageSize, start).RelatedSel("ArticleType").All(&article)
	} else {
		res.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", selectType).All(&article)
		beego.Info("条件筛选数据:", article)
	}

	//分页数据
	obj.Article = article
	return obj
}
func (this *ArticleModel) DelArticle(id int) (res int) {
	o := orm.NewOrm()
	article := ArticleModel{Id: id}
	_, err := o.Delete(&article)
	if err != nil {
		beego.Info("删除失败", err)
		return 500
	}
	return 200
}
func (this *ArticleModel) EditArticle(model ArticleModel) (res int) {
	o := orm.NewOrm()
	_, err := o.Update(&model)
	if err != nil {
		beego.Info("更新失败", err)
		return 500
	}
	return 200
}
func (this *ArticleModel) GetArticleContent(id int) ArticleModel {
	o := orm.NewOrm()
	var article = ArticleModel{Id: id}
	o.Read(&article)
	return article
}

func (this *ArticleTypeModel) AddType(typeName string) int {
	this.TypeName = typeName
	o := orm.NewOrm()
	_, err := o.Insert(this)
	if err != nil {
		beego.Info("插入数据")
		return 500
	}
	return 200
}
func (this *ArticleTypeModel) ShowAddType() []ArticleTypeModel {
	o := orm.NewOrm()
	var resObj []ArticleTypeModel
	_, err := o.QueryTable("ArticleTypeModel").All(&resObj)
	if err != nil {
		beego.Info("数据展示失败", err)
	}
	return resObj
}

func (this *ArticleTypeModel) DelArticleType(id int) int {
	o := orm.NewOrm()
	resObj := ArticleTypeModel{Id: id}
	_, err := o.Delete(&resObj)
	if err != nil {
		beego.Info("删除失败", err)
		return 500
	}
	return 200
}

//查询文章种类
func (this *ArticleTypeModel) GetArticleTypeList() []ArticleTypeModel {
	o := orm.NewOrm()
	var typeObj []ArticleTypeModel
	_, err := o.QueryTable("ArticleTypeModel").All(&typeObj)
	if err != nil {
		beego.Info("文章类型列表查询有误", err)
		return typeObj
	}
	return typeObj
}
