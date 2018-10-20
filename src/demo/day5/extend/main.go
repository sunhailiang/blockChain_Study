package main
import(
	"fmt"
)
//定义父类
type Car struct{
	weight float32
	name string
}
//定义子类
type  Bike struct{
	//继承父类
	Car
	lunzi int
}
//定义指向父类的方法

func(c *Car) Run(){
	fmt.Println("dang~~~");
	
}
 
func main(){
	var b Bike
	//使用父类属性成员
	b.weight=100
	b.name="大卡车"
	//自身的属性成员
	b.lunzi=4
	//使用父类的方法
	b.Run()
	fmt.Println("result",b)
	
}