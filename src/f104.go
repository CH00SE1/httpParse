package main

import (
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

// 启动
func main() {
	for i := 41; i < 61; i++ {
		li5apuu7.ExampleScrape(29, i)
	}
}
