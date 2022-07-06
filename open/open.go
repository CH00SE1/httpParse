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
var tag = 3

const (
	platfrom_paoyou, platfrom_li5apuu7 = "paoyou", "li5apuu7"
)

func main() {
	//hs.Mysql2Redis();
	newPaoYou(1, 21, 10, platfrom_paoyou)
	//newPaoYou(1, 11, 2, platfrom_li5apuu7)
}

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
	for i := num1; i < num2; i++ {
		hs.Paoyou(tag, i)
	}
	defer wg.Done()
}

func THs2(num1, num2 int) {
	for i := num1; i < num2; i++ {
		hs.ExampleScrape(tag, i)
	}
	defer wg.Done()
}

func newPaoYou(num1, num2, size int, funcName string) {
	count := num2 - num1
	if count > 0 {
		wg.Add(size)
		for i := 1; i <= size; i++ {
			if funcName == platfrom_paoyou {
				go THs1(num1+count/size*(i-1), num1+count/size*i)
			}
			if funcName == platfrom_li5apuu7 {
				go THs2(num1+count/size*(i-1), num1+count/size*i)
			}
		}
		// 等待任务全部结束
		wg.Wait()
	} else {
		fmt.Printf("num2 - num1 < 0 , 修改参数")
	}

}
