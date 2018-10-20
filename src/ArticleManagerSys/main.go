package main

import (
	_ "ArticleManagerSys/routers"
	"github.com/astaxie/beego"
	"strconv"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true //启用
	//视图映射
	beego.AddFuncMap("ShowPrePage",ShowPrePage)
	beego.AddFuncMap("ShowNextPage",ShowNextPage)
	beego.Run()
}

//视图映射
//上一页
func ShowPrePage(data int) (string) {
	pageIndex := strconv.Itoa(data - 1)
	return pageIndex
}

//下一页
func ShowNextPage(data int) (string) {
	pageIndex := strconv.Itoa(data + 1)
	return pageIndex
}

