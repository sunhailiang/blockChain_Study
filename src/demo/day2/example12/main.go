package main
import(
	"fmt"
)
func main(){
	var num int;
	fmt.Scanf("%d\n",&num)
	jiecheng(num);
}
func jiecheng(num int) {
	var n int=1;
	var sum int=0;
	for i:=1;i<=num;i++{
		n=n*i;
		sum+=n;
	}
	fmt.Println(sum);
}