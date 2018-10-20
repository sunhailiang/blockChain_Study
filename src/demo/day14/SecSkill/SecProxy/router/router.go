package router

import (
	"github.com/astaxie/beego"
	"demo/day14/SecSkill/SecProxy/controller"
	"github.com/astaxie/beego/logs"
)

//在执行主函数前初始化路由
func init() {
	logs.Debug("router init")
	//匹配指定路由下的controller下的函数  *可以处理各种类型请求，也可改成get或者post
	//秒杀函数
	beego.Router("/seckill",&controller.SkillController{},"*:SecSkill")
	//查询秒杀商品信息
	beego.Router("/secinfo",&controller.SkillController{},"*:SecInfo")
}
