package calc

func Sub(a int, b int,c chan int){
	c<-(a-b)
}