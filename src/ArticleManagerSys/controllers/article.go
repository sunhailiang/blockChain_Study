package controllers

import (
	"github.com/astaxie/beego"
	"ArticleManagerSys/models"
	"path"
	"time"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) ShowAddArticle() {
	var typeObj models.ArticleTypeModel
	var resObj []models.ArticleTypeModel
	resObj = typeObj.GetArticleTypeList()
	if len(resObj) > 0 {
		this.Data["typeList"] = resObj
	} else {
		beego.Info("暂无类型列表数据")
	}
	this.TplName = "add.html"
}

func (this *ArticleController) AddArticle() {
	var article models.ArticleModel
	article.ArticleName = this.GetString("articleName")
	article.Type = this.GetString("selectType")
	article.Content = this.GetString("content")
	f, h, err := this.GetFile("uploadname")
	if article.Type == "" {
		beego.Info("必须选择文章类型")
		return
	}
	defer f.Close()
	ext := path.Ext(h.Filename)
	if (ext != ".jpg" && ext != ".png" && ext != ".gif") || h.Size > 50000 {
		this.Data["upload"] = "图片过大或者图片格式错误"
		return
	}
	if err != nil {
		beego.Info("IMG上传错误")
		return
	}
	var salt = time.Now().Format("2006-01-02 15:04:05")
	article.Img = salt + h.Filename
	article.Count = 0
	//err = this.SaveToFile("uploadname", "./static/userupload/"+article.Img)
	//if err != nil {
	//	beego.Info("图片保存错误", err)
	//	return
	//}
	var articleObj models.ArticleModel
	if articleObj.AddRrticle(article) == 200 {
		this.Redirect("/", 302)
	}
}

//修改
func (this *ArticleController) EditArticle() {
	id, _ := strconv.Atoi(this.GetString("id"))
	article := models.ArticleModel{Id: id}
	article.ArticleName = this.GetString("articleName")
	article.Content = this.GetString("content")
	f, h, err := this.GetFile("uploadname")
	if err != nil {
		beego.Info("不修改图片")
	} else {
		defer f.Close()
		ext := path.Ext(h.Filename)
		if (ext != ".jpg" && ext != ".png" && ext != ".gif") || h.Size > 50000 {
			this.Data["upload"] = "图片过大或者图片格式错误"
			return
		}
		var salt = time.Now().Format("2006-01-02 15:04:05")
		article.Img = salt + h.Filename
		//err = this.SaveToFile("uploadname", ".\\static\\userupload\\"+article.Img)
	}
	article.Count = 0

	if article.EditArticle(article) == 200 {
		this.Redirect("/", 302)
	}
}

//删除
func (this *ArticleController) DelArticle() {
	var article models.ArticleModel
	id, _ := strconv.Atoi(this.GetString("id"))
	if article.DelArticle(id) == 200 {
		this.Redirect("/", 200)
	}

}

//展示内容
func (this *ArticleController) ShowArticleContent() {
	id, _ := strconv.Atoi(this.GetString("id"))
	var article models.ArticleModel
	article = article.GetArticleContent(id)
	this.Data["content"] = article
	this.TplName = "content.html"
}
func (this *ArticleController) ShowEdit() {
	id, _ := strconv.Atoi(this.GetString("id"))
	var article models.ArticleModel
	article = article.GetArticleContent(id)
	this.Data["articleCon"] = article
	this.TplName = "update.html"
}

//添加类型
func (this *ArticleController) AddType() {
	typtName := this.GetString("typeName")
	var addTypeObj models.ArticleTypeModel
	res := addTypeObj.AddType(typtName)
	if res == 200 {
		this.Redirect("/addtype", 302)
	}
}
func (this *ArticleController) ShowAddType() {
	var addTypeObj models.ArticleTypeModel
	var resData = addTypeObj.ShowAddType()
	if len(resData) > 0 {
		this.Data["typeList"] = resData
	}
	beego.Info("没有数据")
	this.TplName = "addType.html"
}
func (this *ArticleController) DelArticleType() {
	var addTypeObj models.ArticleTypeModel
	id, _ := this.GetInt("id")
	if addTypeObj.DelArticleType(id) == 200 {
		beego.Info("删除成功")
	} else {
		beego.Info("删除失败")
	}
	this.Redirect("/addtype", 302)

}
