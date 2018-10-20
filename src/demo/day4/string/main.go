package main
import(
	"fmt"
)
func main(){
	//这里是中文所以占得并不是一个字节，用[]byte会有问题，所以用[]rune x32可以存4个字节来解决
	var str string="我是你爸爸";

	//此处已经转换成数组
	slicestr:=[]rune(str);
    //更改字符
	slicestr[0]='她'
    //重新转换成字符串
	fmt.Println("after",string(slicestr))
}