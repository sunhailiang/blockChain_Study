package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)
type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}
func main()  {
      //打开文件
      file,err:=os.Open("d:/test.txt")
      if err!=nil{
      	fmt.Println("open file err:",err)
	  }
      //操作完成确保文件关闭
      defer file.Close()

      var Count CharCount

      //创建读取的对象
      reder:=bufio.NewReader(file)
      for{
      	str,err:=reder.ReadString('\n')
      	if err==io.EOF{
                break
		}
		if err!=nil{
			fmt.Printf("read file failed, err:%v", err)
			break
		}
		//获取读取行，并用[]rune按照字符切割
		var runeArr=[]rune(str)
		  for _,v:=range runeArr{
			  switch {
			  case v >= 'a' && v <= 'z':
				  fallthrough
			  case v >= 'A' && v <= 'Z':
				  Count.ChCount++
			  case v == ' ' || v == '\t':
				  Count.SpaceCount++
			  case v >= '0' && v <= '9':
				  Count.NumCount++
			  default:
				  Count.OtherCount++
			  }
		  }
	  }
	fmt.Printf("char count:%d\n", Count.ChCount)
	fmt.Printf("num count:%d\n", Count.NumCount)
	fmt.Printf("space count:%d\n", Count.SpaceCount)
	fmt.Printf("other count:%d\n", Count.OtherCount)

}