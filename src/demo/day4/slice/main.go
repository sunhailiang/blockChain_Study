package main
import(
	"fmt"
)
func main(){
	//创建一个数组
	var arr[10]string=[...]string{"A","B","C","D","E","F","G","H","I","J"}
	//数组切片
	var slice=arr[0:7];
	fmt.Println(slice)
	fmt.Println(cap(slice));
	fmt.Println(len(slice));	

}