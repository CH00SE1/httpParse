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
var lock sync.Mutex

// 1907
var tag = 1

const (
	platfrom_paoyou, platfrom_li5apuu7 = "paoyou", "li5apuu7"
)

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

func THs1(num1, num2 int) {
	defer wg.Done()
	for i := num1; i < num2; i++ {
		hs.Paoyou(tag, i)
	}
}

func THs2(num1, num2 int) {
	defer wg.Done()
	for i := num1; i < num2; i++ {
		hs.ExampleScrape(tag, i)
	}
}

func newPaoYou(num1, num2, size int, funcName string) {
	count := num2 - num1
	if count > 0 {
		wg.Add(1)
		for i := 1; i <= size; i++ {
			// 加锁
			lock.Lock()
			if funcName == platfrom_paoyou {
				go THs1(num1+count/size*(i-1), num1+count/size*i)
			}
			if funcName == platfrom_li5apuu7 {
				go THs2(num1+count/size*(i-1), num1+count/size*i)
			}
			// 解锁
			lock.Unlock()
		}
	} else {
		fmt.Printf("num2 - num1 < 0 , 修改参数")
	}
	wg.Wait()
}

func main() {
	syncTpaoyou()
	//newPaoYou(100, 4000, 20, platfrom_paoyou)
	//newPaoYou(1, 1000, 40, platfrom_li5apuu7)
}
