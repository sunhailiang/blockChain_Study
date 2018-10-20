package main
import "fmt"
//定义接口
type Test interface {
	run()
	//repaire()
}
//定义结构体
type Car struct {
	cartype string
	carname string
	weight  float64
}

//实现多态，挂载到指定对象
func (c Car) run() {
	fmt.Printf("一辆载重%v千克的%v类型%v,在路上奔跑", c.weight, c.cartype, c.carname)
}

func main() {
	var b Car = Car{
		carname: "自行车",
		cartype: "人力",
		weight:  100.00,
	}
	var t Test
	t = b
	t.run()
}
