package main

import (
	"dingjing/threadpool"
	"fmt"
	"time"
)

func count(i interface{}) {
	fmt.Println(i)
}

func main() {
	tp := threadpool.New()
	tp.SetFunc(count)
	tp.SetThreadNum(1)
	tp.SetQueueSize(1000)

	for i := 0; i < 10; i++ {
		tp.AddWork(i)
	}

	time.Sleep(time.Second)

	tp.Run()

}
