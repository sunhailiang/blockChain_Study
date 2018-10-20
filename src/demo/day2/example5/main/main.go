package main
import(
	"fmt"
	"time"
)

const(
	man=1
	female=2
)

func main(){
 var now_sec=time.Now().Unix();
 
 if (now_sec % female==0){
	 fmt.Println("gender","female")
 }else{
	fmt.Println("gender","man")	 
 }
	
}