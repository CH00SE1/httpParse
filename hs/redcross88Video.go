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
	"strings"
)

/**
 * @title https://www.redcross88.top/
 * @author xiongshao
 * @date 2022-08-31 16:12:54
 */

const orgRedCross88, RedCross = "https://www.redcross88.top", "红会所"

// 页面请求
func RedCross88Request(classId, page int, className string) {
	oldUrl := orgRedCross88 + "/index.php/vod/type/id/" + strconv.Itoa(classId) + ".html"

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

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	reader, _ := goquery.NewDocumentFromReader(res.Body)

	// 引入数据库连接
	db, _ := db.MysqlConfigure()

	reader.Find("div.colVideoList div.video-elem").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find("a.display").Attr("href")
		photoUrl1, _ := selection.Find("a.display div.img").Attr("style")
		title := selection.Find("a.title").Text()
		replace1 := strings.Replace(photoUrl1, "background-image: url('", "", -1)
		photoUrl := strings.Replace(replace1, "')", "", -1)
		row := redis.KeyExists(title)
		if row != 1 {
			playVideoUrl := orgRedCross88 + href
			m3u8UrlFalse := jyM3u8Url(playVideoUrl)
			m3u8UrlTrue := M3u8Request(m3u8UrlFalse)
			hsInfo := HsInfo{
				Title:    title,
				Url:      playVideoUrl,
				M3u8Url:  m3u8UrlTrue,
				ClassId:  classId,
				PhotoUrl: photoUrl,
				Platform: RedCross + "-" + className,
				Page:     page,
				Location: strconv.Itoa(page) + "-" + strconv.Itoa(i+1),
			}
			marshal, _ := json.Marshal(hsInfo)
			redis.SetKey(title, marshal)
			db.Create(&hsInfo)
		} else {
			PrintfCommon(page, i, href, title, 1, className)
		}
	})

}
