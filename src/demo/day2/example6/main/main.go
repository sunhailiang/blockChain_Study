package main

import(
	"fmt"
	"os"
)
func main(){
	var windowsinfo string =os.Getenv("GOOS")
	var path string=os.Getenv("PATH")
	fmt.Println("windowsinfo:  ",windowsinfo,"PATH: ",path);
}