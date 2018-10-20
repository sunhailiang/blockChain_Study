package main
import (
	"fmt"
)
func changevalue(a *int,b *int){
	temp:=*a;
	*a=*b;
	*b=temp;
	return
}
func main(){
	a:=100;
	b:=200;
	changevalue(&a,&b);
	fmt.Println("a: ",a,"b: ",b)
}