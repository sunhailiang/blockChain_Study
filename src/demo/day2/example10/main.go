package main
import(
	"fmt"
)
func main(){
   
	var n int;
	var m int;
	fmt.Scanf("%d%d",&n,&m)
	fmt.Printf("%d",n)
	//for i:=n;i<m;i++ {
	//	if getprime(i)==true {
	//		fmt.Printf("%d\n",i)
	//	}
	//}

}
//获取素数
//func getprime(n int)bool{
//	//取余为零的，2是最小素数
//      for i:=2;i<n;i++{
//		  if(n%i==0){
//			  return false
//		  }
//	  }
//	  return true
//}