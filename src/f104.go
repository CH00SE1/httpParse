package main

import (
	"fmt"
	"httpParse/li5apuu7"
	"time"
)

// 定时任务
func TimeTask() {
	var ch chan int
	// 定时任务
	ticker := time.NewTicker(time.Second * 30)
	go func() {
		for range ticker.C {
			// 方法
		}
		ch <- 1
	}()
	<-ch
}

func main() {
	for i := 02; i < 61; i++ {
		fmt.Println(i)
		li5apuu7.ExampleScrape(28, i)
	}
}
