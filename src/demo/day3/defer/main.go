package main
import (
	"fmt"
)
func main(){
   var i int;

   defer fmt.Println("这是先进来的",i)
   defer fmt.Println("这是第二个");
   i=100;

   fmt.Println("这是函数执行内容",i)
}
