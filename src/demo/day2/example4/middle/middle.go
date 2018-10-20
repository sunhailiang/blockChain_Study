package middle

import(
	     "fmt"
	_ "demo/day2/example4/last"
)

var Name string
var Age int

func init(){
	Name="this is middle"
	Age=99
	fmt.Println("middleName",Name,"middleAge",Age)
}