package main
import (
	"fmt"
	//起个别名叫add
 	add "demo/day2/example2/add"
)
func main(){
	//使用包别名
	fmt.Println("name",add.Name,"age",add.Age)
}