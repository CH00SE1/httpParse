package hs

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"httpParse/db"
	"httpParse/redis"
	"io"
	"net/http"
	"strconv"
	"strings"
)

/**
 * @title 69yw
 * @author CH00SE1
 * @date 2022-10-22 15:08:35
 */
const Url69yw = "https://69yw5.xyz"
const webName = "69尤物"

func Req69ywHtml() {
	//ParseListVideoInfo(getHtml(url_69yw))
	fmt.Println(ParseClassInfo(Url69yw))
}

// 重新转化请求连接
func IsPage(pageNum int, prefix string) string {
	return Url69yw + prefix + strconv.Itoa(pageNum) + ".html"
}

// 请求页面获取返回
func GetHtml(url string) io.ReadCloser {
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
	return res.Body
}

func ParseClassInfo(url string) map[string]string {
	contentsMap := make(map[string]string)
	html := GetHtml(url)
	reader, _ := goquery.NewDocumentFromReader(html)
	reader.Find("div.nav a").Each(func(i int, selection *goquery.Selection) {
		classId, _ := selection.Attr("href")
		className := selection.Text()
		if className != "精选推荐" {
			contentsMap[classId] = className
		}
	})
	return contentsMap
}

// 分析视频集合页面
func ParseListVideoInfo(body io.ReadCloser, classId, page int, className string) {
	dom, _ := goquery.NewDocumentFromReader(body)
	db, _ := db.MysqlConfigure()
	dom.Find("div.vod").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find("div.vod-img a").Attr("href")
		dataOriginal, _ := selection.Find("div.vod-img a img").Attr("data-original")
		title := selection.Find("div.vod-txt a").Text()
		html := GetHtml(Url69yw + href)
		reader, _ := goquery.NewDocumentFromReader(html)
		videoPlayInfo, _ := reader.Find("div.player iframe").Attr("src")
		index := strings.Index(videoPlayInfo, "?")
		m3u8Url := videoPlayInfo[index+5:]
		//fmt.Println(href, dataOriginal, title, m3u8Url)
		row := redis.KeyExists(title)
		if row != 1 {
			hsInfo := HsInfo{
				Title:    title,
				Url:      Url69yw + href,
				M3u8Url:  m3u8Url,
				PhotoUrl: dataOriginal,
				Platform: webName + className,
				ClassId:  classId,
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
