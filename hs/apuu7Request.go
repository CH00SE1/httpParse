package hs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"httpParse/db"
	"httpParse/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/**
 * @title hs 网页解析
 * @author xiongshao
 * @date 2022-06-22 11:46:35
 */

// 数据保存结构体
type HsInfo struct {
	gorm.Model
	Title    string `gorm:"unique;not null;comment:标题"`
	Url      string
	M3u8Url  string
	Platform string
	ClassId  int
	Page     int
	Location string
}

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

// paoyou视频信息结构体
type PaoYouVideo struct {
	Data struct {
		Id      int    `json:"id"`
		Vid     int    `json:"vid"`
		Pid     int    `json:"pid"`
		Zid     int    `json:"zid"`
		Name    string `json:"name"`
		Playurl string `json:"playurl"`
		Xid     int    `json:"xid"`
		Pay     struct {
			Time   int    `json:"time"`
			Nums   int    `json:"nums"`
			Cion   int    `json:"cion"`
			Msg    string `json:"msg"`
			Btntxt string `json:"btntxt"`
			Init   int    `json:"init"`
		} `json:"pay"`
		Cion    int           `json:"cion"`
		Type    string        `json:"type"`
		Ads     []interface{} `json:"ads"`
		Uvip    int           `json:"uvip"`
		Nexturl string        `json:"nexturl"`
		Vname   string        `json:"vname"`
	} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
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

func ExampleScrape(tag int, page int) (string, int) {
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
	db, _ := db.MysqlConfigure()
	// Find the review items
	doc.Find("div.item a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		//title := s.Find("a").Text()
		title, _ := s.Attr("title")
		href, _ := s.Attr("href")
		url := "http://li5.apuu7.top" + utils.StringStrip(href)
		str += "\"title\":\"" + utils.StringStrip(title) + "\" ,\"url\":\"" + url + "\"},\n"
		//fmt.Println("title:", title, "url:", url)
		row := db.Where("(title) = @title", sql.Named("title", utils.StringStrip(title))).Find(&HsInfo{}).RowsAffected
		if row != 1 {
			if strings.Contains(url, "http://li5.apuu7.top/index.php/vod/play") {
				obj := getM3u8Obj(url)
				m3u8url := M3u8UrlParse(obj)
				//fmt.Println("m3u8Url:", m3u8url)
				// 插入数据
				db.Create(&HsInfo{Title: utils.StringStrip(title), Url: utils.StringStrip(url), M3u8Url: utils.StringStrip(m3u8url), ClassId: tag, Platform: "li5apuu7", Page: page, Location: strconv.Itoa(i) + "-[" + strconv.Itoa(i/6+1) + "," + strconv.Itoa(i%6) + "]"})
			}
		}
	})
	// 每页停止6秒
	//time.Sleep(2 * 1e9)
	return str, page
}

// url https://paoyou.ml
func Paoyou(tag int, page int) {

	initial_url := "https://paoyou.ml/"

	url := initial_url + "lists/" + strconv.Itoa(tag) + "/" + strconv.Itoa(page) + ".html"

	fmt.Println("\n请求 url : ", url)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "paoyou.ml")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("cookie", "Hm_lvt_c0b6c8564ce67088dca63919bcc664b9=1655732450,1655990235,1656234817; Hm_lpvt_c0b6c8564ce67088dca63919bcc664b9=1656234837")
	req.Header.Add("sec-ch-ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "none")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err2 := goquery.NewDocumentFromReader(res.Body)

	if err2 != nil {
		log.Fatal(err2)
	}

	// 引入数据库连接
	db, _ := db.MysqlConfigure()

	doc.Find("ul.fed-list-info li a.fed-list-pics").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		title, _ := s.Attr("title")
		row := db.Where("(title) = @title", sql.Named("title", utils.StringStrip(title))).Find(&HsInfo{}).RowsAffected
		if row != 1 {
			jid := getDataJid(initial_url + href)
			m3u8_url := getM3U8URl(jid)
			// 获取输出
			fmt.Printf("\n[第%d页 第%d个] -> [href:%s , title:%s , m3u8_url:%s]\n", page, i+1, href, title, m3u8_url)
			// 插入数据
			db.Create(&HsInfo{Title: utils.StringStrip(title), Url: utils.StringStrip(initial_url + href), M3u8Url: utils.StringStrip(m3u8_url), ClassId: tag, Platform: "paoyou", Page: page, Location: strconv.Itoa(i) + "-[" + strconv.Itoa(i/6+1) + "," + strconv.Itoa(i%6) + "]"})
		} else {
			fmt.Printf("\n[第%d页 第%d个] -> [href:%s , title:%s , row:%d]\n", page, i+1, href, title, row)
		}

	})

}

// 请求播放页面拿去视频jid
func getDataJid(url string) string {
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "paoyou.ml")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("cookie", "Hm_lvt_c0b6c8564ce67088dca63919bcc664b9=1655732450,1655990235,1656234817; Hm_lpvt_c0b6c8564ce67088dca63919bcc664b9=1656240892")
	req.Header.Add("sec-ch-ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "none")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	doc, err2 := goquery.NewDocumentFromReader(res.Body)

	if err2 != nil {
		log.Fatal(err2)
	}

	// 获取到视频JID
	dataJid, _ := doc.Find("div.video").Attr("data-jid")

	return dataJid
}

// 根据视频jid获取m3u8地址
func getM3U8URl(jid string) string {

	url := "https://paoyou.ml/index.php/ajax/vodurl"
	method := "POST"

	text := "jid=" + jid

	payload := strings.NewReader(text)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "paoyou.ml")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("cookie", "Hm_lvt_c0b6c8564ce67088dca63919bcc664b9=1655732450,1655990235,1656234817; Hm_lpvt_c0b6c8564ce67088dca63919bcc664b9=1656240892")
	req.Header.Add("origin", "https://paoyou.ml")
	req.Header.Add("sec-ch-ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var paoyouvideo PaoYouVideo
	json.Unmarshal(body, &paoyouvideo)
	return paoyouvideo.Data.Playurl
}
