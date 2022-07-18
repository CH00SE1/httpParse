package main

import (
	"fmt"
	"sync"
)

/**
 * @title li5apuu7获取数据
 * @author xiongshao
 * @date 2022-06-27 17:08:58
 */

var wg sync.WaitGroup

func main() {
	for i := 0; i < 2; i++ {
		TestReq(1, 21)
		fmt.Println("--------------------------- 分隔 ---------------------------")
	}
}
