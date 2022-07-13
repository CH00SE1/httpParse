package hs

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"httpParse/db"
	"httpParse/redis"
	"httpParse/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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

// madou视频接口返回对象
type MaDouDao struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Data        []struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Status      int    `json:"status"`
		Thumb       string `json:"thumb"`
		Preview     string `json:"preview"`
		Panorama    string `json:"panorama"`
		Description string `json:"description"`
		VideoUrl    string `json:"video_url"`
		Comefrom    int    `json:"comefrom"`
		Tags        []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tags"`
		TestVideoUrl  string `json:"test_video_url"`
		TrySecond     int    `json:"try_second"`
		IsBloger      int    `json:"is_bloger"`
		IsVip         int    `json:"is_vip"`
		ComefromTitle string `json:"comefrom_title"`
	} `json:"data"`
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

// ------------------------------------------------ li5apuu7 ------------------------------------------------
func ExampleScrape(tag, page int) (string, int) {
	// Request the HTML page.
	// http://li5.apuu7.top/index.php/vod/type/id/20/page/2.html
	url := ""
	if page == 1 {
		url += "http://li5.apuu7.top/index.php/vod/type/id/" + strconv.Itoa(tag) + ".html"
	} else {
		url += "http://li5.apuu7.top/index.php/vod/type/id/" + strconv.Itoa(tag) + "/page/" + strconv.Itoa(page) + ".html"
	}

	fmt.Printf("\n请求 url : %s\n", url)

	res, err := http.Get(url)
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
	redis.InitClient()

	// Find the review items
	doc.Find("div.item a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		//title := s.Find("a").Text()
		title, _ := s.Attr("title")
		href, _ := s.Attr("href")
		url := "http://li5.apuu7.top" + utils.StringStrip(href)
		str += "\"title\":\"" + utils.StringStrip(title) + "\" ,\"url\":\"" + url + "\"},\n"
		//row := db.Where("(title) = @title", sql.Named("title", utils.StringStrip(title))).Find(&HsInfo{}).RowsAffected
		newTitle := utils.StringStrip(title)
		row := redis.KeyExists(newTitle)
		if row != 1 {
			if strings.Contains(url, "http://li5.apuu7.top/index.php/vod/play") {
				obj := getM3u8Obj(url)
				m3u8url := M3u8UrlParse(obj)
				fmt.Printf("\nli5apuu7 [第%d页 第%d个] [href:%s title:%s m3u8_url:%s]\n", page, i+1, href, title, m3u8url)
				hsinfo := HsInfo{Title: utils.StringStrip(title),
					Url:     utils.StringStrip(url),
					M3u8Url: utils.StringStrip(m3u8url),
					ClassId: tag, Platform: "li5apuu7",
					Page:     page,
					Location: "[" + strconv.Itoa(i/6+1) + "," + strconv.Itoa(i%6+1) + "]"}
				marshal, _ := json.Marshal(hsinfo)
				redis.SetKey(newTitle, marshal)
				db.Create(&hsinfo)
			}
		} else {
			fmt.Printf("\nli5apuu7 [第%d页 第%d个] [href:%s title:%s row:%d]\n", page, i+1, href, title, row)
		}
	})
	return str, page
}

// 2.1。同步redis数据 遍历redis数据
func Redis2Mysql() {
	Mysql2Redis()
	keys := redis.GetKeyList()
	mysqlDb, err := db.MysqlConfigure()
	if err != nil {
		fmt.Println("connent datebase err:", err)
	}
	for _, key := range keys {
		values, _ := redis.GetKey(key)
		var hsInfo HsInfo
		json.Unmarshal(utils.String2Bytes(values), &hsInfo)
		mysqlDb.Create(&hsInfo)
	}
}

// 测试案列 mysql数据同步redis
func Mysql2Redis() {
	redis.InitClient()
	db, err := db.MysqlConfigure()
	if err != nil {
		fmt.Println(err)
	}
	var infos []HsInfo
	// 查询数据
	db.Find(&infos)
	for _, info := range infos {
		// 添加序列化后的数据到redis
		marshal, _ := json.Marshal(info)
		redis.SetKey(info.Title, marshal)
	}
}

// ------------------------------------------------ paoyou ------------------------------------------------
func Paoyou(page int, videoName string, map1, map2 map[string]string) {

	url := "https://paoyou.ml"

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
		paoyouDataSave(newTitle, url, href, className, page, i)
	})
}

// <---------------------paoyou数据处理---------------------->
func paoyouDataSave(newTitle, initial_url, href, className string, page, i int) {
	// 引入数据库 mysql + redis
	db, _ := db.MysqlConfigure()
	redis.InitClient()
	row := redis.KeyExists(newTitle)
	if row != 1 {
		jid := getDataJid(initial_url + href)
		m3u8_url := getM3U8URl(jid)
		// 获取输出
		fmt.Printf("\npaoyou [第%d页,第%d个] [href:%s title:%s m3u8_url:%s]\n", page, i+1, href, newTitle, m3u8_url)
		hsinfo := HsInfo{
			Title:    newTitle,
			Url:      utils.StringStrip(initial_url + href),
			M3u8Url:  utils.StringStrip(m3u8_url),
			ClassId:  page,
			Platform: "paoyou" + " -- " + className,
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

// ------------------------------------------------ madou ------------------------------------------------
func MaodouReq(page int) []byte {

	date := strings.Replace(time.Now().Format("2006-01-02"), "-", "", -1)

	hour := time.Now().Hour() - 1

	str_hour := ""

	// 计算时间
	if hour < 10 {
		str_hour += "0" + strconv.Itoa(hour)
	} else {
		str_hour += strconv.Itoa(hour)
	}

	//nnp35.com -- 91tv
	url := "https://nnp35.com/upload_json_live/" + date + "/videolist_" + date + "_" + str_hour + "_2_-_-_100_" + strconv.Itoa(page) + ".json"
	//jsonmdtv.md29.tv -- maodou
	//url := "https://jsonmdtv.md29.tv/upload_json_live/" + date + "/videolist_" + date + "_" + str_hour + "_2_-_-_100_" + strconv.Itoa(page) + ".json"
	//json.wtjfjil.cn
	//url := "https://json.wtjfjil.cn/upload_json_live/" + date + "/videolist_zh-cn_" + date + "_" + str_hour + "_-_-_-_50_" + strconv.Itoa(page) + ".json"
	method := "GET"

	fmt.Printf("\n请求 url : %s\n", url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err, page)
	}
	req.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("Referer", "")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err, page)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected EOF") && len(body) != 0 {
			log.Fatal(err, page)
		}
	}
	return body
}

// 数据转化存储
func DataParseSave(body []byte) {
	var maDouDao MaDouDao
	json.Unmarshal(body, &maDouDao)

	datas := maDouDao.Data
	db, _ := db.MysqlConfigure()
	redis.InitClient()
	for i, data := range datas {
		row := redis.KeyExists(data.Title)
		if row != 1 {
			fmt.Printf("\nmadou [第%d页 第%d个] [href:%s title:%s m3u8_url:%s]\n", maDouDao.CurrentPage, i+1, "https://uh2089he.com"+data.TestVideoUrl, data.Title, strings.Replace(data.VideoUrl, "\\/", "/", -1))
			hsInfo := HsInfo{
				Title:    data.Title,
				Url:      "https://uh2089he.com" + data.TestVideoUrl,
				M3u8Url:  strings.Replace(data.VideoUrl, "\\/", "/", -1),
				ClassId:  maDouDao.CurrentPage,
				Platform: "madou -- " + data.ComefromTitle,
				Page:     maDouDao.CurrentPage,
				Location: "[" + strconv.Itoa((i+1)/6+1) + "," + strconv.Itoa((i+1)%6+1) + "]"}
			marshal, err := json.Marshal(hsInfo)
			if err != nil {
				fmt.Println("hsInfo json 序列化失败")
			}
			redis.SetKey(data.Title, marshal)
			db.Create(&hsInfo)
		} else {
			fmt.Printf("\nmadou [第%d页 第%d个] [href:%s title:%s row:%d]\n", maDouDao.CurrentPage, i+1, "https://uh2089he.com"+data.TestVideoUrl, data.Title, row)
		}
	}
	// 创建文件
	//bytes2String := utils.Bytes2String(body)
	//utils.CreateFile(&bytes2String, "D:\\MadouData\\response\\", "madou_"+
	//time.Now().Format("[2006-01-02-15-04-05]_page_")+strconv.Itoa(maDouDao.PerPage), ".json")
}

// ------------------------------------------------ maomi ------------------------------------------------
func MaomoRequest(page int) {

	initial_url := "https://www.2b3r3.com/"
	// "猫咪推荐"
	// 国产精品 美女主播 短视频 中文字幕
	videoTitle := "美女主播"

	url := newUrl(initial_url, videoTitle, page)

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
	redis.InitClient()

	dom.Find("a.video-pic").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		title, _ := selection.Attr("title")
		row := redis.KeyExists(title)
		if row != 1 {
			m3u8_url := parseMaomiViderPlay(initial_url + href)
			fmt.Printf("herf,title,m3u8_url = %s , %s , %s\n", initial_url+href, title, m3u8_url)
			hsInfo := HsInfo{
				Title:    title,
				Url:      initial_url + href,
				M3u8Url:  m3u8_url,
				ClassId:  page,
				Platform: "maomi -- " + videoTitle,
				Page:     page,
				Location: "[" + strconv.Itoa((i+1)/4+1) + "," + strconv.Itoa((i+1)%4+1) + "]"}
			marshal, _ := json.Marshal(hsInfo)
			redis.SetKey(title, marshal)
			db.Create(&hsInfo)
		} else {
			fmt.Printf("maomi title:%s href:%s row:%d\n", title, href, row)
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
	return "页面解析没有取到地址"
}

// 转为请求url
func newUrl(initial_url, videoTitle string, page int) string {
	if page == 1 {
		return initial_url + "/shipin/list-" + videoTitle + ".html"
	}
	return initial_url + "/shipin/list-" + videoTitle + "-" + strconv.Itoa(page) + ".html"
}
