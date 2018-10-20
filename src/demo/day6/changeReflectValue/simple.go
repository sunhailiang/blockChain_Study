package main
import (
	"reflect"
	"fmt"
)
func main()  {
	var f float64=9.999
	rtest(&f)
	fmt.Println(f)
}
func rtest(obj interface{})  {
	//简单反射赋值
	var v= reflect.ValueOf(obj).Elem()
	v.SetFloat(8.888)

}


