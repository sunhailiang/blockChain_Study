package main
import (
	"fmt"
	"encoding/json"
)
func main(){
   var stu student=student{
	   Name:"harry", 
	   Age:27,
	   Score:99,
   }
   //json序列化
   data,err:=json.Marshal(stu)
   if err!=nil{
	   fmt.Println("json failed",err)
	   return
   }

   fmt.Println("data",string(data));


}

type student struct{
	//json通过读取key,获得字段的说明
	Name string `json:"stu_name"`
	Age int `json:"stu_age"`
	Score int `json:"stu_score"`
}

