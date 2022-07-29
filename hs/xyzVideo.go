package hs

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"httpParse/db"
	"httpParse/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/**
 * @title xyzVideo
 * @author xiongshao
 * @date 2022-07-29 14:13:33
 */

const xyz_url = "https://www.96yz115.xyz"

var purls = [5]string{
	"https://11bfbregwv.com",
	"https://11bfbrzwxr.com",
	"https://11bfbyyhxf.com",
	"https://11bfbfyzqp.com",
	"https://11bfbqnrpz.com"}

func newXyzUrl(page int) string {
	if page == 1 {
		return xyz_url + "/Html/60"
	}
	return xyz_url + "/Html/60/index-" + strconv.Itoa(page) + ".html"
}

func downloandVideo(url string) string {
	html := ".html"
	replace := strings.Replace(url, html, "", -1)
	index := strings.LastIndex(replace, "/")
	return xyz_url + "/Html/player/play-" + replace[index+1:] + "-1-1.html"
}

func videoM3u8(text string) string {
	open := `varplayurl=kele('`
	end := `');`
	index_end := strings.Index(text, end)
	parseUrl := text[len(open):index_end]
	return strings.Replace(parseUrl, "https://d.9xxav.com", purls[1], -1) + ".m3u8"
}

// 请求页面拿到数据
func Xyzrequest(page int) {
	method := "GET"

	url := newXyzUrl(page)

	fmt.Println(url)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	dom, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	db, _ := db.MysqlConfigure()
	dom.Find("div.box ul li").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find("a").Attr("href")
		videoName := selection.Find("h3").Text()
		videoDate := selection.Find("span.movie_date").Text()
		var m3u8Url string
		if strings.Contains(href, ".html") {
			fmt.Println(i, xyz_url+href, utils.StringStrip(videoName), utils.StringStrip(videoDate))
			downloand := downloandVideo(xyz_url + href)
			m3u8Url = XyzRequestVideoPlayInfo(downloand)
		}
		hsinfo := HsInfo{
			Title:    videoName,
			Url:      xyz_url + href,
			M3u8Url:  m3u8Url,
			ClassId:  60,
			Platform: "XyzVideo",
			Page:     page,
			Location: videoDate,
		}
		db.Create(&hsinfo).Callback()
	})

}

// 请求视频视频播放页面
func XyzRequestVideoPlayInfo(url string) string {

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

	dom, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var m3u8Url string

	dom.Find("script").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if i == 12 {
			m3u8Url = videoM3u8(utils.StringStrip(text))
		}
	})

	return m3u8Url

}
