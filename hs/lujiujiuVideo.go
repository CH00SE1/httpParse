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
 * @title 撸久久视频
 * @author xiongshao
 * @date 2022-08-22 14:29:27
 */

const org_url = "https://lu99av1.xyz"

func LujiujiuRequest(classId, page int, className string) {

	oldUrl := org_url + "/index.php/vod/type/id/" + strconv.Itoa(classId) + ".html"

	url := ConvertUrl(oldUrl, page)

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

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	reader, _ := goquery.NewDocumentFromReader(res.Body)

	// 引入数据库连接
	db, _ := db.MysqlConfigure()

	reader.Find("ul.stui-vodlist li.stui-vodlist__item").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find("a").Attr("href")
		title, _ := selection.Find("a").Attr("title")
		photoUrl, _ := selection.Find("a").Attr("data-original")
		row := redis.KeyExists(title)
		if row != 1 {
			fmt.Println(className + " - " + strconv.Itoa(page) + "页")
			playVideoUrl := Display2Video(org_url, href)
			m3u8Url := PlayVideoM3u8Info(playVideoUrl, 8)
			hsInfo := HsInfo{
				Title:    title,
				Url:      playVideoUrl,
				M3u8Url:  m3u8Url,
				ClassId:  classId,
				PhotoUrl: photoUrl,
				Platform: "撸久久-" + className,
				Page:     page,
				Location: strconv.Itoa(page) + "-" + strconv.Itoa(i+1),
			}
			db.Create(&hsInfo)
			marshal, _ := json.Marshal(hsInfo)
			redis.SetKey(title, marshal)
		} else {
			PrintfCommon(page, i, href, title, 1, className)
		}
	})

}
