package main
import "fmt"
//定义接口规范
type Cars interface {
	 run()
	 getCarType()
	 didi()
}
//定义Car类
 type Car struct {
 	 name string
 	 cartype string
 	 takeweight float64
 }
 //实现接口
func (c *Car)run()  {
    fmt.Printf("一辆载重%f的%s类型的%s在高速狂奔\n",c.takeweight,c.cartype,c.name)
}
func (c *Car)getCarType()  {
	fmt.Printf("车型：%s \n",c.cartype)
}
func (c *Car)didi()  {
    fmt.Printf("%s车，在鸣笛\n",c.name)
}
func main()  {
	var  bus Car=Car{
		name:"大巴",
		cartype:"长途大巴",
		takeweight:1444.22,
	}
	var getcar Cars=&bus
	getcar.run()
	getcar.getCarType()
	getcar.didi()
}
