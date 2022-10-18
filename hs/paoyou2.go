package hs

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"httpParse/db"
	"httpParse/redis"
	"net/http"
	"strconv"
)

/**
 * @title paoyou2
 * @author CH00SE1
 * @date 2022-10-18 16:16:12
 */

// 请求页面获取视频信息
func RequestPaoyou2(pageNum int, classId string) {
	center, pageObeject, html := "/index.php/vod/show/id/", "/page/", ".html"
	var url string
	if pageNum == 1 {
		url = paoyou_url + center + classId + html
	} else {
		url = paoyou_url + center + classId + pageObeject + strconv.Itoa(pageNum) + html
	}
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	dom, _ := goquery.NewDocumentFromReader(res.Body)
	db, _ := db.MysqlConfigure()
	dom.Find("li.iidx-vodlist__item a.iidx-vodlist__thumb").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		dataOriginal, _ := selection.Attr("data-original")
		title, _ := selection.Attr("title")
		row := redis.KeyExists(title)
		if row != 1 {
			script := requestPlayVideoPage(paoyou_url + href)
			parse := scriptInfoParse(script)
			hsInfo := paoyouDataSave(parse, pageNum, i, title, dataOriginal)
			db.Create(&hsInfo).Callback()
		} else {
			PrintfCommon(pageNum, i, href, title, row, "paoyou*")
		}
	})
}
