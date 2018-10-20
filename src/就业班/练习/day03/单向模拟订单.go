package main

import "fmt"

type orderInfo struct {
	id int
}

func main() {

	ch:=make(chan orderInfo)
	go producer(ch)

	customer(ch)
}
//生产者
func producer (orderChan chan<- orderInfo){
	for i:=0;i<100;i++ {
		order:=orderInfo{id:i}
		orderChan<-order
	}
	close(orderChan)//通知写入完成
}
//消费者
func customer(order <-chan orderInfo)  {
    for v:=range order{//for 取出
    	fmt.Println("订单号为：",v)
	}
}