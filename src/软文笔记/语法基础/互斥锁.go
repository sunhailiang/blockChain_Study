package main
//互斥锁：传统并发编程对共享资源进行访问控制的主要手段，它由标准库sync中的Mutex结构体类型表示。
//两个指针方法：ync.Mutex类型只有两个公开的指针方法，Lock和Unlock。Lock锁定当前的共享资源，Unlock进行解锁。

//线程同步：互斥量，同一时刻只有一个线程持有该锁，锁一旦被占用其他线程阻塞，直到解锁
//func main() {
//	go person1("hello")
//	go person2("word")
//
//	for {
//		;
//	}
//}
//
//var mutex sync.Mutex
//
//func printer(str string) {
//	//加锁
//	mutex.Lock()  //上了互斥锁之后顺序执行
//	for _, v := range str {
//		fmt.Printf(string(v))
//		time.Sleep(time.Millisecond * 400)
//	}
//	mutex.Unlock()//干掉了锁就没了顺序，因为cpu分配时间片就随机输出了
//
//}
//
//func person1(str string) {
//	printer(str)
//}
//func person2(str string) {
//	printer(str)
//}
