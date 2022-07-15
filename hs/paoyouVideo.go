package hs

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"httpParse/db"
	"httpParse/redis"
	"httpParse/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/**
 * @title paoyou视频
 * @author xiongshao
 * @date 2022-07-14 15:43:22
 */

// 主页
const paoyou_url = "https://paoyou.ml"

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

// ------------------------------------------------ paoyou ------------------------------------------------
func Paoyou(page int, videoName string, map1, map2 map[string]string) {

	newUrl, className := PaoyouNewUrl(videoName, page, map1, map2)

	fmt.Printf("\nurl:%s\tvideoName:%s\n", newUrl, className)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, newUrl, nil)

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

	dom, err2 := goquery.NewDocumentFromReader(res.Body)

	if err2 != nil {
		log.Fatal(err2)
	}
	dom.Find("ul.fed-list-info li a.fed-list-pics").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		title, _ := s.Attr("title")
		newTitle := utils.StringStrip(title)
		paoyouDataSave(newTitle, paoyou_url, href, className, page, i)
	})
}

// paoyou数据处理
func paoyouDataSave(newTitle, url, href, className string, page, i int) {
	// 引入数据库 mysql + redis
	db, _ := db.MysqlConfigure()
	redis.InitClient()
	row := redis.KeyExists(newTitle)
	if row != 1 {
		jid := getDataJid(url + href)
		m3u8_url := getM3U8URl(jid)
		// 获取输出
		fmt.Printf("\npaoyou [第%d页,第%d个] [href:%s title:%s m3u8_url:%s]\n", page, i+1, href, newTitle, m3u8_url)
		hsinfo := HsInfo{
			Title:    newTitle,
			Url:      utils.StringStrip(url + href),
			M3u8Url:  utils.StringStrip(m3u8_url),
			ClassId:  page,
			Platform: "paoyou*" + className,
			Page:     page,
			Location: strconv.Itoa(i + 1)}
		marshal, _ := json.Marshal(hsinfo)
		redis.SetKey(newTitle, marshal)
		db.Create(&hsinfo)
	} else {
		fmt.Printf("\npaoyou [第%d页,第%d个] [href:%s title:%s row:%d]\n", page, i+1, href, newTitle, row)
	}
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

// 查询平台分类
func PaoyouFindClass() (map[string]string, map[string]string) {

	url := paoyou_url + "/lists/1/1.html"

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
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	dom, err2 := goquery.NewDocumentFromReader(res.Body)

	if err2 != nil {
		log.Fatal(err2)
	}
	paoyouMap1 := make(map[string]string)
	paoyouMap2 := make(map[string]string)
	dom.Find("dl").First().Next().Find("dd a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		text := selection.Text()
		paoyouMap1[text] = href
	})
	dom.Find("li.fed-col-sm2 a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		text := selection.Text()
		paoyouMap2[text] = href
	})
	return paoyouMap1, paoyouMap2
}

// 拼接请求 url
// classId 1:基础分类 2.专项分类
func PaoyouNewUrl(classname string, page int, map1, map2 map[string]string) (string, string) {
	html := ".html"
	url1, ok1 := map1[classname]
	url2, ok2 := map2[classname]
	if ok1 {
		if page == 1 {
			return paoyou_url + url1, classname
		}
		return paoyou_url + strings.Replace(url1, html, "", -1) + "/page/" + strconv.Itoa(page) + html, classname
	}
	if ok2 {
		replace := strings.Replace(url2, html, "", -1)
		index := strings.LastIndex(replace, "/")
		return paoyou_url + replace[:index+1] + strconv.Itoa(page) + html, classname
	}
	return "*", classname
}
