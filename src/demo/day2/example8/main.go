package main
import (
	"fmt"
	"math/rand"
)
func main(){
	for i:=0;i<10;i++{
		fmt.Println("int",rand.Int())
	}
	for i:=0;i<10;i++{
		fmt.Println("<100",rand.Intn(100));
	}
	for i:=0;i<10;i++{
		fmt.Println("float",rand.Float32());
	}
}