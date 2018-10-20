package main
import(
	"fmt"
)
func main(){
	var arr1[5]string=[...]string{"a","b","c","d","e"};
   
	slice:=arr1[:];	
	//未扩容之前还指向原数组的内存地址
	slice[0]="w";
    fmt.Println("before arr1:",arr1)
     aaaa:=[]string{"v","j"};
	//切片是可变的，但容量超过了原数组最大值，此处发生了另外开辟内存
	slice=append(slice,aaaa);
	slice=append(slice,"g");
	slice=append(slice,"h");
	slice=append(slice,"i");
	slice=append(slice,"j");
    //此处修改已经与原数组不相关
	slice[1]="ssss";
	// fmt.Println("after arr1:",arr1);
	fmt.Println("slice:",slice);
	fmt.Printf("len[%d],cap[%d]",len(slice),cap(slice));


}