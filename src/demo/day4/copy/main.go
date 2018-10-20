package main
import(
	"fmt"
)
func main(){
	var slice[]int=[]int{1,2,3,4,5,6,7,8}
	fmt.Println("原:",slice);
	var cslice=make([]int,5)
	//复制
	copy(cslice,slice)
	fmt.Println("拷贝",cslice)
}