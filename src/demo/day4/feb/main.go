package main
import(
	"fmt"
)
func main(){
  for i:=0;i<10;i++{
	  res:=feb(i);
      fmt.Println(res);
	}

}
func feb(n int)int{
    if n<=1{
		return 1
	} 
	return feb(n-2)+feb(n-1);
}
