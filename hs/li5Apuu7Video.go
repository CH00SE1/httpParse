package hs

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"httpParse/db"
	"httpParse/redis"
	"httpParse/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/**
 * @title li5.apuu7 视频
 * @author xiongshao
 * @date 2022-06-22 11:46:35
 */

const li5Apuu7_url = "http://li5.apuu7.top/index.php"

// javascript对象
type Player_aaaa struct {
	gorm.Model
	Flag     string `json:"flag"`
	Encrypt  int    `json:"encrypt"`
	Trysee   int    `json:"trysee"`
	Points   int    `json:"points"`
	Link     string `json:"link"`
	LinkNext string `json:"link_next"`
	LinkPre  string `json:"link_pre"`
	Url      string `json:"url"`
	UrlNext  string `json:"url_next"`
	From     string `json:"from"`
	Server   string `json:"server"`
	Note     string `json:"note"`
	Id       string `json:"id"`
	Sid      int    `json:"sid"`
	Nid      int    `json:"nid"`
}

// 请求每个视频链接拿到m3u8下载地址对象部分
func getM3u8Obj(url string) string {
	get, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer get.Body.Close()
	if get.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", get.StatusCode, get.Status)
	}
	reader, err2 := goquery.NewDocumentFromReader(get.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	html := reader.Find("div.pl-l script").Text()
	return html
}

// 目前网页解析获取方法
func M3u8UrlParse(url string) string {
	// 1.把\/转为/
	str1 := strings.Replace(url, "\\/", "/", -1)
	// 2.获取=后面部分
	index := strings.Index(str1, "=")
	str2 := str1[index+1:]
	// 3.string转为结构体
	var player_aaaa Player_aaaa
	json.Unmarshal([]byte(str2), &player_aaaa)
	return player_aaaa.Url
}

// ------------------------------------------------ li5apuu7 ------------------------------------------------
func ExampleScrape(tag, page int) {

	url := ""
	if page == 1 {
		url += li5Apuu7_url + "/vod/type/id/" + strconv.Itoa(tag) + ".html"
	} else {
		url += li5Apuu7_url + "/vod/type/id/" + strconv.Itoa(tag) + "/page/" + strconv.Itoa(page) + ".html"
	}

	fmt.Printf("\n请求 url : %s\n", url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 引入数据库连接
	db, _ := db.MysqlConfigure()

	// Find the review items
	doc.Find("div.item a").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Attr("title")
		href, _ := s.Attr("href")
		photo_url, _ := s.Find("div.img img.thumb").Attr("data-original")
		url := "http://li5.apuu7.top" + utils.StringStrip(href)
		newTitle := utils.StringStrip(title)
		row := redis.KeyExists(newTitle)
		if row != 1 {
			if strings.Contains(url, li5Apuu7_url+"/vod/play") {
				obj := getM3u8Obj(url)
				m3u8url := M3u8UrlParse(obj)
				hsinfo := HsInfo{
					Title:   utils.StringStrip(title),
					Url:     utils.StringStrip(url),
					M3u8Url: utils.StringStrip(m3u8url),
					ClassId: tag, Platform: "li5apuu7",
					PhotoUrl: photo_url,
					Page:     page,
					Location: "[" + strconv.Itoa(i/6+1) + "," + strconv.Itoa(i%6+1) + "]",
				}
				marshal, _ := json.Marshal(hsinfo)
				redis.SetKey(newTitle, marshal)
				db.Create(&hsinfo)
			}
		} else {
			PrintfCommon(page, i+1, utils.StringStrip(url), title, row, "li5apuu7")
		}
	})
}
