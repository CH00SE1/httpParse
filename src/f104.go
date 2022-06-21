package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"httpParse/model/ths"
	"httpParse/utils"
	"log"
	"net/http"
	"strconv"
)

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
		fmt.Printf("Review %d: %s\n", i, utils.StringStrip(href))
		str += "\"title\":\"" + utils.StringStrip(title) + "\" ,\"url\":\"" + "http://li5.apuu7.top" + utils.StringStrip(href) + "\"},\n"
		// 插入数据
		db.Create(&ths.THs{Title: utils.StringStrip(title), Url: "http://li5.apuu7.top" + utils.StringStrip(href)})
	})
	return str
}

func main() {

}
