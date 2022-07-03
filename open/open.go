package main

import (
	"httpParse/hs"
	"sync"
)

/**
 * @title tmp
 * @author xiongshao
 * @date 2022-06-30 17:03:29
 */

// 全局变量
var wg sync.WaitGroup
var tag = 2

func flush() {
	defer wg.Done()
	for i := 130; i < 200; i++ {
		hs.Paoyou(tag, i)
	}
}

func syncTpaoyou() {
	wg.Add(1)
	go flush()
	wg.Wait()
}

func main() {
	syncTpaoyou()
}
