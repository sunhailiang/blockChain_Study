package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func  main()  {
	 inputfile:="d:/readorwrite/read.txt"
	 outfile:="d:/readorwrite/write.txt"
	 buf,err:=ioutil.ReadFile(inputfile)
	 if err!=nil{
	 	fmt.Fprintf(os.Stderr,"file err:%s",err)
	 	return
	 }
	 fmt.Printf("%s \n",string(buf))
	 err=ioutil.WriteFile(outfile,buf,0x64)
	if err!=nil {
		panic(err.Error())
	}
}
