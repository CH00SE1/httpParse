package main

import (
	"fmt"
	"httpParse/hs"
	"sync"
)

/**
 * @title li5apuu7获取数据
 * @author xiongshao
 * @date 2022-06-27 17:08:58
 */

var wg sync.WaitGroup

func TE1(tag int) {
	defer wg.Done()
	for i := 400; i < 500; i++ {
		hs.ExampleScrape(tag, i)
	}
}

func TE2(tag int) {
	defer wg.Done()
	for i := 240; i < 260; i++ {
		hs.ExampleScrape(tag, i)
	}
}

func TE3(tag int) {
	defer wg.Done()
	for i := 260; i < 300; i++ {
		hs.ExampleScrape(tag, i)
	}
}

func flush(tag int) {
	defer wg.Done()
	for i := 2; i < 20; i++ {
		hs.ExampleScrape(tag, i)
	}
}

func test() {
	tag := 28
	wg.Add(1)
	go flush(tag)
	wg.Wait()
}
func nancy(num1, num2, size int) {
	var ints []int
	for i := 1; i <= size; i++ {
		ints = append(ints, num1+(num2-num1)/size*i)
		fmt.Printf("[%d,%d]\n", num1+(num2-num1)/size*(i-1), num1+(num2-num1)/size*i)
	}
	fmt.Println(ints)
}

func main() {
	nancy(219, 1000, 20)
}
