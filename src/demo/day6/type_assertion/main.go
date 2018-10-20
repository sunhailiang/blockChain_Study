package main

import "fmt"

func main()  {
	Assertions(22,"dddd",true)
}

//自定义类型
type person struct {

}

//批量断言 Assertion
func Assertions(items...interface{})  {
	for index,value:=range items{
		switch value.(type) {//判断值的类型
		case bool:
			fmt.Printf("%d param is bool,the value is %v \n",index,value)
		case int,int32,int64:
			fmt.Printf("%d param is int,the value is %v \n",index,value)
		case float32,float64:
			fmt.Printf("%d param is float,the value is %v \n",index,value)
		case string:
			fmt.Printf("%d param is string,the value is %v \n",index,value)
		case byte:
			fmt.Printf("%d param is byte,the value is %v \n",index,value)
		case person://判断自定义类型
			fmt.Printf("%d param is person,the value is %v \n",index,value)
		}
	}
}

//单个断言
func assertion(obj interface{}){
	value,ok:=obj.(int)
	//如果失败ok的值就是false
	if ok==false{
		fmt.Println("conver failed")
		return
	}
	//如果成功拿到对象的值
	fmt.Printf("the value is %d\n",value)
}

