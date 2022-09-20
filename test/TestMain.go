package main

import (
	"sync"
)

/**
 * @title 测试主方法
 * @author xiongshao
 * @date 2022-06-27 17:08:58
 */

var wg sync.WaitGroup

func main() {
	//// 测试方法 10-AV中文视频 3-经典国产 2-国产传媒
	//for i := 1; i < 20; i++ {
	//	hs.JinyuislandRequest(2, i, "国产传媒")
	//}
	//for i := 1; i < 100; i++ {
	//	hs.RedCross88Request(91, i, "经典国产")
	//}
	TestReq(100, 200)
}
