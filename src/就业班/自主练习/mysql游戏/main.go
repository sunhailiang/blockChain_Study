package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"就业班/自主练习/mysql游戏/pk"
)

func main() {
   db,err:=sql.Open("mysql","root:123456@tcp(192.168.35.41:3306)/pk")
   defer db.Close()
   err=db.Ping()
	if err!=nil {
		fmt.Println("数据库链接失败~")
		log.Fatal(err)
	}else{
		fmt.Println("数据库连接成功~")
	}
	num:=pk.HeroCount(db)
	var h pk.Hero
	h.InitHero(1,db)
	fmt.Println("hero num",num)
}
