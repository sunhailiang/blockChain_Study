package main
import(
	"fmt"
)

//声明结构体
type Person struct{
	 gender string
	 age int
	 weight float32
}
func main(){
	//初始化
	var p Person
	p.gender="男"
	p.age=27
	p.weight=135.6
  
	fmt.Println(p)

}