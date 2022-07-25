package open

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
var tag = 20

const (
	Platfrom_paoyou, Platfrom_li5apuu7, Platfrom_madou, Platfrom_maomi, Platfrom_G, Platfrom_cape = "paoyou", "li5apuu7", "madou", "maomi", "G.", "cape"
	className                                                                                     = "日本"
)

// 多线程方法
func GetHs(num1, num2, size int, funcName string) {
	count := num2 - num1
	if count > 0 {
		wg.Add(size)
		for i := 1; i <= size; i++ {
			n1, n2 := num1+count/size*(i-1), num1+count/size*i
			if funcName == Platfrom_paoyou {
				go THs1(n1, n2)
			}
			if funcName == Platfrom_li5apuu7 {
				go THs2(n1, n2)
			}
			if funcName == Platfrom_madou {
				go THs3(n1, n2)
			}
			if funcName == Platfrom_maomi {
				go THs4(n1, n2)
			}
			if funcName == Platfrom_G {
				go THs5(n1, n2)
			}
			if funcName == Platfrom_cape {
				go THs6(n1, n2)
			}
		}
		// 等待任务全部结束
		wg.Wait()
	} else {
		fmt.Printf("num2 - num1 < 0 , 修改参数")
	}
}

// <----------------------------------------- Paoyou ----------------------------------------->
func THs1(num1, num2 int) {
	//map1 := hs.PaoyouFindClass()
	//for i := num1; i < num2; i++ {
	//	hs.Paoyou(i, className, map1)
	//}
	//defer wg.Done()
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
		maDouDao, Type, urlType := new(hs.Org).MaodouReq(i)
		hs.DataParseSave(maDouDao, Type, urlType)
	}
	defer wg.Done()
}

// <----------------------------------------- maomi ----------------------------------------->
func THs4(num1, num2 int) {
	for i := num1; i <= num2; i++ {
		new(hs.Org).MaomiRequest(i)
	}
	defer wg.Done()
}

// <----------------------------------------- G. ----------------------------------------->
func THs5(num1, num2 int) {
	for i := num1; i < num2; i++ {
		hs.GRequest(i)
	}
	defer wg.Done()
}

// <----------------------------------------- cape ----------------------------------------->
func THs6(num1, num2 int) {
	for i := num1; i < num2; i++ {
		new(hs.Org).RequestPageInfo(i)
	}
	defer wg.Done()
}
