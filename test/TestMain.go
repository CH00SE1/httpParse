package main

import (
	"fmt"
	"httpParse/hs"
	"sync"
)

/**
 * @title li5apuu7获取数据
 * @author xiongshao
 * @date 2022-06-27 17:08:58
 */

var wg sync.WaitGroup

func flush(tag int) {
	defer wg.Done()
	for i := 2; i < 20; i++ {
		hs.ExampleScrape(tag, i)
	}
}

func main() {
	for i := 0; i < 2; i++ {
		TestReq(100, 200)
		fmt.Println("--------------------------- 分隔 ---------------------------")
	}
}
