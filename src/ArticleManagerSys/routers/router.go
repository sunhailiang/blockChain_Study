package routers

import (
	"ArticleManagerSys/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.AuthController{}, "get:LoginView;post:Login")
	beego.Router("/regist", &controllers.AuthController{}, "get:RegistView;post:Regist")
	beego.Router("/logout", &controllers.AuthController{}, "get:Logout")
	beego.Router("/addarticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:AddArticle")
	beego.Router("/delete", &controllers.ArticleController{}, "get:DelArticle")
	beego.Router("/edit", &controllers.ArticleController{}, "post:EditArticle;get:ShowEdit")
	beego.Router("/content", &controllers.ArticleController{}, "get:ShowArticleContent")
	beego.Router("/addtype", &controllers.ArticleController{}, "get:ShowAddType;post:AddType")
	beego.Router("/deltype", &controllers.ArticleController{}, "get:DelArticleType")
}
