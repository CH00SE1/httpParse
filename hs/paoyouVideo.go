package hs

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"httpParse/db"
	"httpParse/redis"
	"httpParse/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
 * @title paoyou视频
 * @author xiongshao
 * @date 2022-07-14 15:43:22
 */

// 主页
const paoyou_url = "https://paoyou.ml"

// paoyou aaa返回对象
type PaoyouDao struct {
	Flag     string `json:"flag"`
	Encrypt  int    `json:"encrypt"`
	Trysee   int    `json:"trysee"`
	Points   int    `json:"points"`
	Link     string `json:"link"`
	LinkNext string `json:"link_next"`
	LinkPre  string `json:"link_pre"`
	VodData  struct {
		VodName     string `json:"vod_name"`
		VodActor    string `json:"vod_actor"`
		VodDirector string `json:"vod_director"`
		VodClass    string `json:"vod_class"`
	} `json:"vod_data"`
	Url     string `json:"url"`
	UrlNext string `json:"url_next"`
	From    string `json:"from"`
	Server  string `json:"server"`
	Note    string `json:"note"`
	Id      string `json:"id"`
	Sid     int    `json:"sid"`
	Nid     int    `json:"nid"`
}

// ------------------------------------------------ paoyou ------------------------------------------------
// 1.拼接请求 url
func PaoyouNewUrl(classname string, page int, map1 map[string]string) (string, string) {
	if page == 1 {
		url := map1[classname]
		return paoyou_url + url, classname
	} else {
		html := ".html"
		url := map1[classname]
		index := strings.Index(url, html)
		newUrl := url[:index]
		return paoyou_url + newUrl + "/page/" + strconv.Itoa(page) + html, classname
	}
	return "*", classname
}

// 2.请求页码
func Paoyou(page int, videoName string, map1 map[string]string) {

	newUrl, className := PaoyouNewUrl(videoName, page, map1)

	fmt.Printf("\nurl:%s\tvideoName:%s\n", newUrl, className)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, newUrl, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "paoyou.ml")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("cookie", "Hm_lvt_1f12b0865d866ae1b93514870d93ce89=1655378802; Hm_lvt_da72459bf70a79f74af84e92497546d0=1658243321; Hm_lvt_c0b6c8564ce67088dca63919bcc664b9=1658028529,1658243253,1658305701,1658306244; Hm_lpvt_c0b6c8564ce67088dca63919bcc664b9=1658309647")
	req.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "none")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62")

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
	redis.InitClient()
	db, _ := db.MysqlConfigure()
	var HsInfos []*HsInfo
	dom.Find("ul.stui-vodlist li.stui-vodlist__item a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		text1, _ := s.Attr("title")
		text := utils.StringStrip(text1)
		row := redis.KeyExists(text)
		if row != 1 {
			videoPage := requestPlayVideoPage(paoyou_url + href)
			parse := scriptInfoParse(videoPage)
			hsInfo := paoyouDataSave(parse, page, i, text)
			HsInfos = append(HsInfos, &hsInfo)
		} else {
			PrintfCommon(page, i, href, text, row, "paoyou*"+videoName)
		}
	})
	db.Table("t_hs_info3").CreateInBatches(HsInfos, 500).Callback()
	time.Sleep(8 * time.Second)
}

// 3.请求播放页面 拿到script标签var player_aaaa={} json对象
func requestPlayVideoPage(url string) string {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "paoyou.ml")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("cookie", "Hm_lvt_1f12b0865d866ae1b93514870d93ce89=1655378802; Hm_lvt_c0b6c8564ce67088dca63919bcc664b9=1658028529,1658243253,1658305701,1658306244; Hm_lpvt_c0b6c8564ce67088dca63919bcc664b9=1658310215; Hm_lvt_da72459bf70a79f74af84e92497546d0=1658243321,1658310216; Hm_lpvt_da72459bf70a79f74af84e92497546d0=1658310216")
	req.Header.Add("referer", "https://paoyou.ml/index.php/vod/type/id/1.html")
	req.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62")

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

	// 拿到三层div第一个scrpit标签
	text := dom.Find("div.stui-pannel div.stui-player div.stui-player__video script").First().Text()

	return text
}

// 4.转义一下返回数据
func scriptInfoParse(text string) PaoyouDao {
	// 1.把\/转为/
	str1 := strings.Replace(text, "\\/", "/", -1)
	// 2.获取=后面部分
	index := strings.Index(str1, "=")
	str2 := str1[index+1:]
	var paoyoudao PaoyouDao
	json.Unmarshal([]byte(str2), &paoyoudao)
	return paoyoudao
}

// 5.paoyou数据处理
func paoyouDataSave(paoyouDao PaoyouDao, page, i int, title string) HsInfo {
	play_url := paoyou_url + paoyouDao.Link
	hsInfo := HsInfo{
		Title:    title,
		Url:      play_url,
		M3u8Url:  paoyouDao.Url,
		ClassId:  page,
		Platform: "paoyou*" + paoyouDao.VodData.VodClass,
		Page:     page,
		Location: "[" + strconv.Itoa((i+1)/4+1) + "," + strconv.Itoa((i+1)%4+1) + "]",
	}
	marshal, _ := json.Marshal(hsInfo)
	redis.SetKey(title, marshal)
	return hsInfo
}

// 查询平台分类
func PaoyouFindClass() (map[string]string, []string) {

	url := paoyou_url

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
	paoyouMap := make(map[string]string)
	var paoyouArray []string
	dom.Find("div.stui-header__menu ul.clearfix li a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		text := selection.Text()
		stripText := utils.StringStrip(text)
		paoyouMap[text] = href
		if !strings.Contains(stripText, "首") && !strings.Contains(stripText, "更") {
			paoyouArray = append(paoyouArray, stripText)
		}
	})
	dom.Find("div.stui-pannel__head a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		text := selection.Text()
		stripText := utils.StringStrip(text)
		paoyouMap[text] = href
		if !strings.Contains(stripText, "首") && !strings.Contains(stripText, "更") {
			paoyouArray = append(paoyouArray, stripText)
		}
	})
	return paoyouMap, paoyouArray
}
