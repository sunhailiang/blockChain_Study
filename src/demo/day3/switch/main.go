package main
import(
	"fmt"
	"math/rand"
)

func main(){
	var num int=rand.Intn(100);
		
	for{
		var numbr int
		fmt.Scanf("%d\n",&numbr);
		 if checkNum(numbr,num){
			 break;
		 }
	}
	
}

func checkNum(n int,num int)bool{
		var flag bool=false;
		switch {
		case n==num:
			fmt.Println("猜对了");
			flag=true
		case num>n:
			fmt.Println("大于：",n)
		case num<n:
			fmt.Println("小于:",n)
		}	
       return flag
	
}

