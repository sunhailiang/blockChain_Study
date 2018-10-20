package main

import (
	"os"
	"fmt"
)

func main() {
	 fs,err:= os.Create("./textDoc/a.txt")
	 if err!=nil{
	 	fmt.Println("create failed")
	 }
	 str:="床前明月光，疑是地上霜"

	 b:=[]byte(str)

	 fs.Write(b)

	 defer fs.Close()

}
