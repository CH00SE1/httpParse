package src

import (
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
