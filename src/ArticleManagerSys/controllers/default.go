package controllers

import (
	"github.com/astaxie/beego"
	"ArticleManagerSys/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var articleModel models.ArticleModel
	var resObj models.ResPageObj
	if c.GetSession("user") == nil {
		c.Redirect("/login", 302)
	}
	index, err := c.GetInt("pageIndex")
	if err != nil {
		index = 1
	}
	articleModel.Type = c.GetString("selectType")
	beego.Info("index", index)
	beego.Info("resObj.Page.TotalPage", resObj.Page.TotalPage)
	resObj = articleModel.GetArticleList(index, articleModel.Type)
	c.Data["articleList"] = resObj.Article
	c.Data["pageInfo"] = resObj.Page
	c.Data["articleTypes"] = resObj.Type
	c.TplName = "index.html"
}
