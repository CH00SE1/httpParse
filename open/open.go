package main

import (
	"fmt"
	"httpParse/hs"
	"sync"
)

/**
 * @title paoyou tmp
 * @author xiongshao
 * @date 2022-06-30 17:03:29
 */

// 全局变量
var wg sync.WaitGroup

// 1907
var tag = 3

func flush() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		hs.Paoyou(tag, i)
	}
}

func syncTpaoyou() {
	wg.Add(1)
	go flush()
	wg.Wait()
}

func THs(num1, num2 int) {
	defer wg.Done()
	for i := num1; i < num2; i++ {
		hs.Paoyou(tag, i)
	}
}

func newPaoYou(num1, num2, size int) {
	if num2-num1 > 0 {
		wg.Add(1)
		for i := 1; i <= size; i++ {
			go THs(num1+(num2-num1)/size*(i-1), num1+(num2-num1)/size*i)
		}
	} else {
		fmt.Printf("num2 - num1 > 0 , 修改参数")
	}
	wg.Wait()
}

func main() {
	newPaoYou(1, 11, 5)
}
