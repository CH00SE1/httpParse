package hs

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	paoyou_url = "https://paoyou.ml"
)

/**
 * @title 查询平台分类
 * @author xiongshao
 * @date 2022-07-11 09:12:57
 */

// 查询平台分类
func PaoyouFindClass() (map[string]string, map[string]string) {

	url := paoyou_url + "/lists/1/1.html"

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "paoyou.ml")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("cookie", "Hm_lvt_c0b6c8564ce67088dca63919bcc664b9=1655732450,1655990235,1656234817; Hm_lpvt_c0b6c8564ce67088dca63919bcc664b9=1656234837")
	req.Header.Add("sec-ch-ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "none")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	dom, err2 := goquery.NewDocumentFromReader(res.Body)

	if err2 != nil {
		log.Fatal(err2)
	}
	paoyouMap1 := make(map[string]string)
	paoyouMap2 := make(map[string]string)
	dom.Find("dl").First().Next().Find("dd a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		text := selection.Text()
		paoyouMap1[text] = href
	})
	dom.Find("li.fed-col-sm2 a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		text := selection.Text()
		paoyouMap2[text] = href
	})
	return paoyouMap1, paoyouMap2
}

// 拼接请求 url
// classId 1:基础分类 2.专项分类
func PaoyouNewUrl(classname string, page int, map1, map2 map[string]string) (string, string) {
	html := ".html"
	url1, ok1 := map1[classname]
	url2, ok2 := map2[classname]
	if ok1 {
		if page == 1 {
			return paoyou_url + url1, classname
		}
		return paoyou_url + strings.Replace(url1, html, "", -1) + "/page/" + strconv.Itoa(page) + html, classname
	}
	if ok2 {
		replace := strings.Replace(url2, html, "", -1)
		index := strings.LastIndex(replace, "/")
		return paoyou_url + replace[:index+1] + strconv.Itoa(page) + html, classname
	}
	return "*", classname
}
