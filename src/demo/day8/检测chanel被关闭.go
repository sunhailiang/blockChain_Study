package main
import "fmt"
func main()  {
   ch:=make(chan int,10)
	for i:=0;i< 10;i++{
		ch<-i
	}
	//关闭chanel
	close(ch)
	//方法一
	for {
		var b int
		//ok判断是否输出结束
		b,ok:=<-ch
		//结束即跳出循环
		if !ok{
			fmt.Println("chanel is close")
			break
		}
		fmt.Println(b)
	}
	//for rang//遍历长度，满足即推出
	for v:=range ch{
		 fmt.Println(v)
	}
}
