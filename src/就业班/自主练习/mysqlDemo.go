package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.35.117:3306)/huoying")
	if err != nil {
		fmt.Println("sql open err:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("db.Ping err:", err)
		return
	} else {
		fmt.Println("ping success")
	}
	var heroName string
	var heroId int
	err = db.QueryRow("select heroName,heroId from hero where heroName='漩涡鸣人';").Scan(&heroName, &heroId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(heroName, heroId)

}
