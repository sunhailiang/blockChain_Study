package main
import (
	"encoding/json"
	"fmt"
)
type user struct {
	Username string
	Nickname string
	Age int
	Phone int
	Gender string
}
func main(){
	var user user=user{
	   Username:"王小明",
	   Nickname:"齐天大圣",
	   Age:18,
	   Phone:110,
	   Gender:"男",
	}
	//json序列化
	 data,err:=json.Marshal(user)
    if err!=nil{
    	fmt.Println("json.Marshal failed err:",err)
    	return
	}
	//序列化成功输出结果
	fmt.Println("json:",string(data))

}

