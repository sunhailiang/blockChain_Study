package main
import(
	"fmt"
)
func main(){
	//简易声明方式
	m:=make(map[int]string,10)
	m[1]="aaaa"
	m[2]="bbbb"
	m[3]="cccc"
	fmt.Println("map",m)
}