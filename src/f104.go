package main

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

func ExampleScrape(page int) string {
	// Request the HTML page.
	res, err := http.Get("http://li5.apuu7.top/index.php/vod/type/id/30/page/" + strconv.Itoa(page) + ".html")
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
			get, err1 := http.Get(url)
			if err1 != nil {
				log.Fatal(err1)
			}
			defer get.Body.Close()
			if get.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}
			reader, err2 := goquery.NewDocumentFromReader(get.Body)
			if err2 != nil {
				log.Fatal(err2)
			}
			html := reader.Find("div.pl-l script").Text()
			m3u8url := M3u8UrlParse(html)
			//fmt.Println("m3u8Url:", m3u8url)
			// 插入数据
			db.Create(&ths.THs{Title: utils.StringStrip(title), Url: utils.StringStrip(url), M3u8Url: utils.StringStrip(m3u8url)})
		}

	})
	return str
}

// 定时任务
func TimeTask() {
	var ch chan int
	// 定时任务
	ticker := time.NewTicker(time.Second * 30)
	go func() {
		for range ticker.C {
			ExampleScrape(2)
		}
		ch <- 1
	}()
	<-ch
}

func main() {
	for i := 3; i < 30; i++ {
		ExampleScrape(i)
	}
}
