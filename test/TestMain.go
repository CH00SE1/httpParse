package main

import (
	"httpParse/hs"
	"sync"
)

/**
 * @title li5apuu7获取数据
 * @author xiongshao
 * @date 2022-06-27 17:08:58
 */

var wg sync.WaitGroup

func TE1(tag int) {
	for i := 20; i < 100; i++ {
		hs.ExampleScrape(tag, i)
	}
}

func TE2(tag int) {
	for i := 100; i < 150; i++ {
		hs.ExampleScrape(tag, i)
	}
}

func TE3(tag int) {
	for i := 150; i < 200; i++ {
		hs.ExampleScrape(tag, i)
	}
}

func main() {
	wg.Add(1)
	TE1(32)
	TE2(32)
	TE3(32)
	wg.Wait()
}
