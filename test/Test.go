package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

/**
 * @title 模拟请求方法
 * @author xiongshao
 * @date 2022-07-05 08:48:51
 */

func req(num int) string {
	defer wg.Done()
	url := "http://localhost:8080/m3u8/list/" + strconv.Itoa(num)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "请求错误"
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "请求错误"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "解析错误"
	}
	fmt.Printf("[%d] --> res : {test}\n", num)
	return string(body)
}

func TestReq() {
	for i := 1000; i < 11200; i++ {
		wg.Add(1)
		go req(i)
	}
	wg.Wait()
}
