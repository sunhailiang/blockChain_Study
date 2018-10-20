package main
import(
	"fmt"
)

//自定义类型
type cal func(int,int) int;

//加法
func add(a int,b int) int{
	return a+b;
}
//减法
func sub (a int,b int) int{
	return a-b;
}
//除法
func  div(a int,b int) int{
	return int(a/b);
}
//乘法
func mult(a int,b int) int{
	return a*b;
}

//统一计算函数,此刻调用自定义类型 cal
func calc(c cal,a,b int ){
	res:=c(a,b);
     fmt.Println("结果",res);
}

func main(){
	var num1,num2 int;
	var ty string;
	fmt.Scanf("%s\n%d\n%d",&ty,&num1,&num2)
	fmt.Println("运算符号",ty)
	switch ty{
	case "+":calc(add,num1,num2)
	case "-":calc(sub,num1,num2)
	case "*":calc(mult,num1,num2)
	case "/":calc(div,num1,num2)
	} 
}