package pk

import (
	"database/sql"
	"fmt"
)

type Hero struct {
	HeroId     int    //英雄编号
	HeroName   string //英雄姓名
	HeroATK    int    //英雄基础攻击力
	HeroDEF    int    //英雄基础防御力
	HeroHP     int    //英雄基础血量
	HeroAddATK int    //英雄每级追加攻击力
	HeroAddDEF int    //英雄每级追加防御力
	HeroAddHP  int    //英雄每级追加血量
	HeroInfo   string //英雄简介
}

func HeroCount(db *sql.DB) int {
	var count int
	err := db.QueryRow("select count(*) from hero;").Scan(&count)
	if err != nil {
		fmt.Println("select error")
		return -1
	}
	return count
}

func (this *Hero) InitHero(id int,db *sql.DB) {
	//预编译
	stmt, _ := db.Prepare("select heroId,heroName,heroATK,heroDEF,heroHP,heroAddATK,heroAddDef,heroAddHp,heroInfo FROM hero WHERE heroId = ?;")
	err := stmt.QueryRow(id).Scan(&this.HeroId, &this.HeroName, &this.HeroATK, &this.HeroDEF, &this.HeroAddHP, &this.HeroATK, &this.HeroAddDEF, &this.HeroHP, &this.HeroInfo)
	if err != nil {
		fmt.Println("select data error\n", err)
		return
	}
	fmt.Println(this.HeroId, this.HeroName, this.HeroATK, this.HeroDEF, this.HeroAddHP, this.HeroATK, this.HeroAddDEF, this.HeroHP, this.HeroInfo)
}
