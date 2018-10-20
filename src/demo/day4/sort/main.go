package main
import(
	"fmt"
	"sort"
)
func main(){
	//字符串数组排序
	var str =[...]string{"b","d","w","a","j","l"};
	//因为数组本身是个值类型，徐艳转换成切片灵活排序
	 sort.Strings(str[:]);
	 fmt.Println("string: ",str);
	 
	 //整型排序
	 var ints=[...]int{11,6,3,8,5,55,7}
	 sort.Ints(ints[:])
	 fmt.Println("int: ",ints)
	 
	 //浮点型排序

	 var floats=[...]float64{1.2,3.3,4.2,5.6,2.1};
	 sort.Float64s(floats[:])
	 fmt.Println("float: ",floats)


	 //查找---获取排序后的下标位置
	  var indexs =[...]string{"b","d","w","a","j","l"};
	  //需要先转换为切片
	  res:=sort.SearchStrings(indexs[:],"d");
	  fmt.Println(res);

}