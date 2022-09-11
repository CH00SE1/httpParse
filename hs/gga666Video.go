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
 * @title 玖爱视频
 * @author xiongshao
 * @date 2022-08-18 14:31:06
 */

const org_g66_url = "https://gga996.com"

// 请求数据
func Gga666Request(classId, page int, className string) {

	oldUrl := org_g66_url + "/index.php/vod/type/id/" + strconv.Itoa(classId) + ".html"

	url := ConvertUrl(oldUrl, page)

	fmt.Println(url)

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

	reader.Find("ul.thumbnail-group li").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find("a.thumbnail").Attr("href")
		photoUrl, _ := selection.Find("a.thumbnail img").Attr("data-original")
		title := selection.Find("div.video-info a").Text()
		location := selection.Find("div.video-info p").Text()
		row := redis.KeyExists(title)
		if row != 1 {
			playUrl := Display2Video(org_g66_url, href)
			m3u8Url := PlayVideoM3u8Info(playUrl, 8)
			hsInfo := HsInfo{
				Title:    title,
				Url:      playUrl,
				M3u8Url:  m3u8Url,
				ClassId:  classId,
				PhotoUrl: photoUrl,
				Platform: "玖爱视频-" + className,
				Page:     page,
				Location: "[" + strconv.Itoa(i/4+1) + "," + strconv.Itoa(i%4+1) + "]观看信息:" + location,
			}
			marshal, _ := json.Marshal(hsInfo)
			redis.SetKey(title, marshal)
			db.Create(&hsInfo)
		} else {
			PrintfCommon(page, i, href, title, 1, className)
		}
	})

}
