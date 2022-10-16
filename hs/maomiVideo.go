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
 * @title maomi视频
 * @author xiongshao
 * @date 2022-07-14 15:42:12
 */

const (
	maomi_url                                             = "https://www.b3b5t.com"
	guochanjingpin, meinvzhubo, duanshipin, zhongwenzhimu = "国产精品", "美女主播", "短视频", "中文字幕"
)

// ------------------------------------------------ maomi ------------------------------------------------
// 旧版本
func (org Org) MaomiRequest(page int) {

	videoTitle := guochanjingpin

	url := convertUrl(maomi_url, videoTitle, page)

	method := "GET"

	fmt.Printf("url:%s\n", url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "www.70a89b4819be.com")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("cookie", "sessionid=c143c61f-aa50-42bd-8aea-5f8e871945e9; Hm_lvt_c4994262310cf443b674a94adc2b0319=1657278367; Hm_lvt_2c2eaee7858675aced3fad3d524be9bb=1657278367; _ga=GA1.2.2023943553.1657278368; _gid=GA1.2.720095917.1657278368; Hm_lpvt_2c2eaee7858675aced3fad3d524be9bb=1657278464; Hm_lpvt_c4994262310cf443b674a94adc2b0319=1657278464; _gat_gtag_UA_207595667_3=1; playss=5")
	req.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.49")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	dom, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	db, _ := db.MysqlConfigure()

	dom.Find("a.video-pic").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		title, _ := selection.Attr("title")
		row := redis.KeyExists(title)
		if row != 1 {
			m3u8_url := parseMaomiViderPlay(maomi_url + href)
			hsInfo := HsInfo{
				Title:    title,
				Url:      maomi_url + href,
				M3u8Url:  m3u8_url,
				ClassId:  page,
				Platform: "maomi*" + videoTitle,
				Page:     page,
				Location: "[" + strconv.Itoa((i+1)/4+1) + "," + strconv.Itoa((i+1)%4+1) + "]"}
			marshal, _ := json.Marshal(hsInfo)
			redis.SetKey(title, marshal)
			db.Create(&hsInfo)
		} else {
			PrintfCommon(page, i+1, title, href, row, "maomi"+videoTitle)
		}
	})
}

// 新文件
func (new New) MaomiRequest(page int, videoTitle string) {

	url := convertUrl(maomi_url, videoTitle, page)

	method := "GET"

	fmt.Printf("url:%s\n", url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "www.70a89b4819be.com")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("cookie", "sessionid=c143c61f-aa50-42bd-8aea-5f8e871945e9; Hm_lvt_c4994262310cf443b674a94adc2b0319=1657278367; Hm_lvt_2c2eaee7858675aced3fad3d524be9bb=1657278367; _ga=GA1.2.2023943553.1657278368; _gid=GA1.2.720095917.1657278368; Hm_lpvt_2c2eaee7858675aced3fad3d524be9bb=1657278464; Hm_lpvt_c4994262310cf443b674a94adc2b0319=1657278464; _gat_gtag_UA_207595667_3=1; playss=5")
	req.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.49")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	dom, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	db, _ := db.MysqlConfigure()

	dom.Find("ul.content-list li.content-item").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find("a").Attr("href")
		title, _ := selection.Find("a").Attr("title")
		PhotoUrl, _ := selection.Find("img").Attr("data-original")
		selection.Find("")
		row := redis.KeyExists(title)
		if row != 1 {
			m3u8_url := parseMaomiViderPlay(maomi_url + href)
			hsInfo := HsInfo{
				Title:    title,
				Url:      maomi_url + href,
				M3u8Url:  m3u8_url,
				ClassId:  page,
				Platform: "maomi*" + videoTitle,
				Page:     page,
				PhotoUrl: PhotoUrl,
				Location: strconv.Itoa(page) + "-" + strconv.Itoa(i+1),
			}
			db.Create(&hsInfo)
			marshal, _ := json.Marshal(hsInfo)
			redis.SetKey(title, marshal)
		} else {
			PrintfCommon(page, i+1, title, href, row, "maomi"+videoTitle)
		}
	})
}

// 解析播放页面拿到视频下载地址
func parseMaomiViderPlay(play_url string) string {

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, play_url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Cookie", "sessionid=810fa458-918f-47cc-aec8-6856c1df2377")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	dom, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 获取第一个script标签节点的值
	first := dom.Find("script").First().Text()
	s := strings.Replace(first, " ", "", -1)
	replace := strings.Replace(s, "\n", "", -1)
	if replace != "" {
		sIndex, cIndex, lIndex := "varvideo='", "';varm3u8_host='", "';varm3u8_host1='"
		startIndex := strings.Index(replace, sIndex)
		centerIndex := strings.Index(replace, cIndex)
		lastIndex := strings.Index(replace, lIndex)
		return replace[centerIndex+len(cIndex):lastIndex] + replace[startIndex+len(sIndex):centerIndex]
	}
	return "*"
}

// 根据页码转化请求url
func convertUrl(url, videoTitle string, page int) string {
	if page == 1 {
		return url + "/shipin/list-" + videoTitle + ".html"
	}
	return url + "/shipin/list-" + videoTitle + "-" + strconv.Itoa(page) + ".html"
}
