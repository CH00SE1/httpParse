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
	platfrom_paoyou, platfrom_li5apuu7, platfrom_madou, platfrom_maomi = "paoyou", "li5apuu7", "madou", "maomi"
	videoName                                                          = "string"
)

func main() {
	//hs.Mysql2Redis()
	//getHs(1, 501, 10, platfrom_paoyou)
	//getHs(300, 400, 5, platfrom_li5apuu7)
	//getHs(1, 61, 5, platfrom_madou)
	getHs(76, 534, 2, platfrom_maomi)
}

// <----------------------------------------- Paoyou ----------------------------------------->
func THs1(num1, num2 int) {
	for i := num1; i < num2; i++ {
		hs.Paoyou(tag, i, videoName)
	}
	defer wg.Done()
}

// <----------------------------------------- li5apuu7 ----------------------------------------->
func THs2(num1, num2 int) {
	for i := num1; i < num2; i++ {
		hs.ExampleScrape(tag, i)
	}
	defer wg.Done()
}

// <----------------------------------------- madou ----------------------------------------->
func THs3(num1, num2 int) {
	for i := num1; i < num2; i++ {
		maDouDao := hs.MaodouReq(i)
		hs.DataParseSave(maDouDao)
	}
	defer wg.Done()
}

// <----------------------------------------- maomi ----------------------------------------->
func THs4(num1, num2 int) {
	for i := num1; i <= num2; i++ {
		hs.MaomoRequest(i)
	}
	defer wg.Done()
}

// 多线程方法
func getHs(num1, num2, size int, funcName string) {
	count := num2 - num1
	if count > 0 {
		wg.Add(size)
		for i := 1; i <= size; i++ {
			n1, n2 := num1+count/size*(i-1), num1+count/size*i
			if funcName == platfrom_paoyou {
				go THs1(n1, n2)
			}
			if funcName == platfrom_li5apuu7 {
				go THs2(n1, n2)
			}
			if funcName == platfrom_madou {
				go THs3(n1, n2)
			}
			if funcName == platfrom_maomi {
				go THs4(n1, n2)
			}
		}
		// 等待任务全部结束
		wg.Wait()
	} else {
		fmt.Printf("num2 - num1 < 0 , 修改参数")
	}

}
