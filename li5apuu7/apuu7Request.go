package li5apuu7

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"httpParse/model/ths"
	"httpParse/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
 * @title li5apuu7 网页解析
 * @author xiongshao
 * @date 2022-06-22 11:46:35
 */

// javascript对象
type Player_aaaa struct {
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

func ExampleScrape(tag int, page int) string {
	// Request the HTML page.
	// http://li5.apuu7.top/index.php/vod/type/id/20/page/2.html
	res, err := http.Get("http://li5.apuu7.top/index.php/vod/type/id/" + strconv.Itoa(tag) + "/page/" + strconv.Itoa(page) + ".html")
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

	str := ""

	// 引入数据库连接
	db, _ := utils.DB()

	// Find the review items
	doc.Find("div.item a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		//title := s.Find("a").Text()
		title, _ := s.Attr("title")
		href, _ := s.Attr("href")
		url := "http://li5.apuu7.top" + utils.StringStrip(href)
		str += "\"title\":\"" + utils.StringStrip(title) + "\" ,\"url\":\"" + url + "\"},\n"
		//fmt.Println("title:", title, "url:", url)
		if strings.Contains(url, "http://li5.apuu7.top/index.php/vod/play") {
			obj := getM3u8Obj(url)
			m3u8url := M3u8UrlParse(obj)
			//fmt.Println("m3u8Url:", m3u8url)
			// 插入数据
			db.Create(&ths.THs{Title: utils.StringStrip(title), Url: utils.StringStrip(url), M3u8Url: utils.StringStrip(m3u8url)})
		}
	})
	// 每页停止6秒
	time.Sleep(5 * 1e9)
	return str
}
