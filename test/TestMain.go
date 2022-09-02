package main

import (
	"httpParse/hs"
	"sync"
)

/**
 * @title 测试主方法
 * @author xiongshao
 * @date 2022-06-27 17:08:58
 */

var wg sync.WaitGroup

func main() {

	// 测试方法
	//for i := 1; i < 100; i++ {
	//	hs.JinyuislandRequest(10, i, "AV中文视频")
	//}
	for i := 100; i < 200; i++ {
		hs.RedCross88Request(91, i, "经典国产")
	}

}
