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
	org_tyms_url, m3u8_tyms_prefix = "https://tyms74.xyz", "https://xiusebf2.com"
)

// 请求视频页面
func Tyms74Request(classId, page int, className string) {

	oldUrl := org_tyms_url + "/index.php/vod/type/id/" + strconv.Itoa(classId) + ".html"

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
	redis.InitClient()

	reader.Find("li a.vodlist_thumb").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		title, _ := selection.Attr("title")
		photoUrl, _ := selection.Attr("data-original")
		row := redis.KeyExists(title)
		if row != 1 {
			playVideoUrl := Display2Video(org_tyms_url, href)
			m3u8UrlFalse := PlayVideoM3u8Info(playVideoUrl, 9)
			m3u8UrlTrue := m3u8Request(m3u8UrlFalse)
			hsInfo := HsInfo{
				Title:    title,
				Url:      playVideoUrl,
				M3u8Url:  m3u8UrlTrue,
				ClassId:  classId,
				PhotoUrl: photoUrl,
				Platform: "桃颜蜜色" + className,
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
func m3u8Request(url string) string {

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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	s := string(body)

	if !strings.Contains(s, "/") {
		return s
	}

	index := strings.Index(s, "/")

	s2 := s[index:]

	return m3u8_tyms_prefix + utils.StringStrip(s2)

}
