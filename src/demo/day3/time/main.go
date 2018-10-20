package main
import(
	"time"
	"fmt"
)
func main(){
   //开始时间
   star:=time.Now().UnixNano();
   //花费时间
   test();
   end:=time.Now().UnixNano();
   fmt.Println("代码执行时间 ",(end-star)/1000)
}
func test(){
	time.Sleep(time.Millisecond*100);
}