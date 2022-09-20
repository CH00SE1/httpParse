package hs

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"httpParse/db"
	"httpParse/redis"
	"log"
	"net/http"
	"strconv"
)

/**
 * @title jinrijp 今日精品
 * @author xiongshao
 * @date 2022-08-24 08:37:46
 */

const org_jinrijp_url = "https://www.jinrijp.top"

// 请求视频页面
func JinrijpRequest(classId, page int, className string) {

	oldUrl := org_jinrijp_url + "/index.php/vod/type/id/" + strconv.Itoa(classId) + ".html"

	url := ConvertUrl(oldUrl, page)

	fmt.Println("\n" + url)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android............ecko) Chrome/92.0.4515.105 HuaweiBrowser/12.0.4.300 Mobile Safari/537.36")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	reader, _ := goquery.NewDocumentFromReader(res.Body)

	// 引入数据库连接
	db, _ := db.MysqlConfigure()

	reader.Find("div.remove-18 div.col-sm-4").Each(func(i int, selection *goquery.Selection) {
		photoUrl, _ := selection.Find("img").Attr("src")
		href, _ := selection.Find("a").Attr("href")
		title, _ := selection.Find("a").Attr("title")
		row := redis.KeyExists(title)
		if row != 1 {
			fmt.Println(className + " - " + strconv.Itoa(page) + "页")
			playVideoUrl := Display2Video(org_jinrijp_url, href)
			m3u8Url := PlayVideoM3u8Info(playVideoUrl, 7)
			hsInfo := HsInfo{
				Title:    title,
				Url:      playVideoUrl,
				M3u8Url:  m3u8Url,
				ClassId:  classId,
				PhotoUrl: photoUrl,
				Platform: "今日精品-" + className,
				Page:     page,
				Location: "[" + strconv.Itoa(i/4+1) + "," + strconv.Itoa(i%4+1) + "]",
			}
			marshal, _ := json.Marshal(hsInfo)
			redis.SetKey(title, marshal)
			db.Create(&hsInfo)
		} else {
			PrintfCommon(page, i, href, title, 1, className)
		}
	})

}
