package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * @title 模拟请求方法
 * @author xiongshao
 * @date 2022-07-05 08:48:51
 */

func req(num int) string {
	defer wg.Done()

	//url := "http://localhost:8520/m3u8/list/" + strconv.Itoa(num)
	url := "http://localhost:8520/sentinel_client_sale/sales/shop/day/32"

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
	fmt.Printf("[%d] --> res : {%s}\n", num, "123")
	return string(body)
}

func TestReq(num1, num2 int) {
	for i := num1; i < num2; i++ {
		wg.Add(num2 - num1)
		go req(i)
	}
	wg.Wait()
}
