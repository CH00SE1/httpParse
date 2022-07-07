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
var tag = 2

const (
	platfrom_paoyou, platfrom_li5apuu7 = "paoyou", "li5apuu7"
)

func main() {
	//for i := 1; i <= 60; i++ {
	//	maDouDao := hs.MaodouReq(i)
	//	hs.DataParseSave(maDouDao)
	//}
	//hs.Mysql2Redis();
	getHs(600, 700, 10, platfrom_paoyou)
	//getHs(1, 31, 3, platfrom_li5apuu7)
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

func getHs(num1, num2, size int, funcName string) {
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
