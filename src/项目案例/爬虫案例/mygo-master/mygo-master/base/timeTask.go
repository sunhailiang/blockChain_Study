package main

import(
	"time"
)

func WriteWork(writereload func()) {
	go func() {
		writereload()
		for {
			now := time.Now()
			// 下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			defer LogFile.Close()
			writereload()
			<-t.C
		}
	}()
}

func WriteReload(){
	//init work
	fmt.Println("start...")
}

func main(){
	WriteWork(WriteReload)
}
