package main
import(
	"fmt"
)
func main(){
	var num int;
	var num2 int;
   fmt.Scanf("%d,%d\n",&num,&num2)
   for i:=num;i<num2;i++{
	   isNumber(i);
   }
   
}
func isNumber(num int){
	var a,b,c int;
	a=num%10;
	b=(num/10)%10;
	c=(num/100)%10;
	sum:=a*a*a+b*b*b+c*c*c;
	if(sum==num){
		fmt.Printf("%d\n",num)
	}
}
