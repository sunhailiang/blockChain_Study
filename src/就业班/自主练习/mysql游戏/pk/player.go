package pk

import (
	"database/sql"
	"fmt"
)

type Player struct {
	HeroId    int    //英雄Id
	HeroName  string //英雄姓名
	WeaponId  int    //武器Id
	HeroLevel int    //英雄等级
	HeroExp   int    //英雄经验
	Money     int    //英雄所持金钱
	//HeroWeapon *Weapon      //具体的武器指针，用于后期设计武器类记录玩家所持武器

	HeroBaseATK int //英雄基础攻击力
	HeroBaseDEF int //英雄基础防御力
	HeroBaseHP  int //英雄基础血量

	HeroAddATK int //英雄每级追加攻击力
	HeroAddDEF int //英雄每级追加防御力
	HeroAddHP  int //英雄每级追加血量

	HeroCurrentATK int //玩家当前攻击力
	HeroCurrentDEF int //玩家当前防御力
	HeroCurrentHP  int //玩家当前血量
}
//初始化玩家
func (this *Player)InitPlayer(hero Hero,db *sql.DB)  {
	this.HeroId = hero.HeroId
	this.HeroName = hero.HeroName
	this.WeaponId = 0
	this.HeroLevel = 0
	this.HeroExp = 0
	this.Money = 0
	//this.HeroWeapon = nil   //初始化玩家武器
	//清空数据
	_,err:=db.Exec("truncate table player")
	if err!=nil {
		fmt.Println("truncate table error:%v\n",err)
		return
	}
}
