package main

import (
	"sync"
)

/**
 * @title 测试接口
 * @author xiongshao
 * @date 2022-06-27 17:08:58
 */

var wg sync.WaitGroup

func main() {
	for i := 0; i < 2; i++ {
		TestReq(1, 51)
	}
}
