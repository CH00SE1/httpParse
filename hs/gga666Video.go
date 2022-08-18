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
 * @title ggg666Video.go
 * @author xiongshao
 * @date 2022-08-18 14:31:06
 */

type Ggg666 struct {
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
}

const org_g66_url = "https://gga996.com"

// 请求数据
func Gga666Request(classId, page int, className string) {

	oldUrl := org_g66_url + "/index.php/vod/type/id/" + strconv.Itoa(classId) + ".html"

	url := setUrl(oldUrl, page)

	fmt.Println(url)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
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

	reader.Find("ul.thumbnail-group li").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find("a.thumbnail").Attr("href")
		photoUrl, _ := selection.Find("a.thumbnail img").Attr("data-original")
		title := selection.Find("div.video-info").Text()
		row := redis.KeyExists(title)
		if row != 1 {
			playUrl := display2play(href)
			m3u8Url := playVideoM3u8Info(playUrl)
			hsInfo := HsInfo{
				Title:    title,
				Url:      playUrl,
				M3u8Url:  m3u8Url,
				ClassId:  classId,
				PhotoUrl: photoUrl,
				Platform: "玖爱视频-" + className,
				Page:     page,
				Location: "[" + strconv.Itoa(i/4+1) + "," + strconv.Itoa(i%4+1) + "]",
			}
			marshal, _ := json.Marshal(hsInfo)
			redis.SetKey(title, marshal)
			db.Create(&hsInfo)
		} else {
			fmt.Println(title)
		}
	})

}

// 根据页面转换请求url
func setUrl(url string, page int) string {
	html := ".html"
	if page == 1 {
		return url
	}
	index := strings.Index(url, html)
	return url[:index] + "/page/" + strconv.Itoa(page) + html
}

// 转化为播放url
func display2play(url string) string {
	//展示页面
	//https://gga996.com/index.php/vod/detail/id/5792.html
	//博凡页面
	//https://gga996.com/index.php/vod/play/id/5792/sid/1/nid/1.html
	html := ".html"
	index := strings.LastIndex(url, "/")
	s1 := url[index:]
	s2 := strings.Replace(s1, html, "", -1)
	return org_g66_url + "/index.php/vod/play/id" + s2 + "/sid/1/nid/1" + html
}

// 请求播放页面拿到m3u8url
func playVideoM3u8Info(url string) string {

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	reader, _ := goquery.NewDocumentFromReader(res.Body)

	var m3u8_url string

	reader.Find("script").Each(func(i int, selection *goquery.Selection) {
		title := selection.Text()
		if i == 8 {
			s2 := strings.Replace(title, "\\/", "/", -1)
			index := strings.Index(s2, "=")
			s3 := s2[index+1:]
			var ggg666 Ggg666
			json.Unmarshal([]byte(s3), &ggg666)
			m3u8_url = ggg666.Url
		}
	})
	return m3u8_url
}
