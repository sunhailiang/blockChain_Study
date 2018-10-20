package main

import "github.com/gomodule/redigo/redis"

func main() {
	conn,_:=redis.Dial("tcp","192.168.233.149::6379")
	defer  conn.Close()
	conn.Do("set","every",100000)
}
