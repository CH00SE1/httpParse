package hs

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"httpParse/db"
	"httpParse/redis"
	"httpParse/utils"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/**
 * @title tyms74.xyz 桃颜蜜色
 * @author xiongshao
 * @date 2022-08-29 10:05:51
 */

const (
	orgTymsUrl = "https://tyms81.xyz"
)

// 请求视频页面
func Tyms74Request(classId, page int, className string) {

	oldUrl := orgTymsUrl + "/index.php/vod/type/id/" + strconv.Itoa(classId) + ".html"

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

	reader.Find("li a.vodlist_thumb").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		title, _ := selection.Attr("title")
		photoUrl, _ := selection.Attr("data-original")
		row := redis.KeyExists(title)
		if row != 1 {
			playVideoUrl := Display2Video(orgTymsUrl, href)
			m3u8UrlFalse := PlayVideoM3u8Info(playVideoUrl, 9)
			m3u8UrlTrue := M3u8Request(m3u8UrlFalse)
			hsInfo := HsInfo{
				Title:    title,
				Url:      playVideoUrl,
				M3u8Url:  m3u8UrlTrue,
				ClassId:  classId,
				PhotoUrl: photoUrl,
				Platform: "桃颜蜜色-" + className,
				Page:     page,
				Location: strconv.Itoa(i),
			}
			marshal, _ := json.Marshal(hsInfo)
			redis.SetKey(title, marshal)
			db.Create(&hsInfo)
		} else {
			PrintfCommon(page, i, href, title, 1, className)
		}
	})

}

// 拿到真实的m3u8的下载地址
func M3u8Request(url1 string) string {

	url := utils.StringStrip(url1)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android............ecko) Chrome/92.0.4515.105 HuaweiBrowser/12.0.4.300 Mobile Safari/537.36")

	res, err := client.Do(req)

	if res == nil {
		return "nil pointer"
	}

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	s := string(body)

	if strings.Contains(s, ".ts") {
		return url
	}

	if !strings.Contains(s, "/") {
		return s
	}

	if strings.Contains(s, ".m3u8") {

		index := strings.Index(s, "/")

		s2 := s[index:]

		if strings.Contains(url, "/2") {

			urlPrefix := url[:strings.Index(url, "/2")]

			return urlPrefix + utils.StringStrip(s2)

		} else {

			urlPrefix := url[:strings.Index(url, "/")]

			return urlPrefix + utils.StringStrip(s2)

		}

	}

	return "nil"

}
