package hs

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"httpParse/db"
	"httpParse/redis"
	"httpParse/utils"
	"io"
	"net/http"
	"strconv"
	"strings"
)

/**
 * @title 公共模块
 * @author xiongshao
 * @date 2022-07-14 15:46:21
 */

// 数据保存结构体t_hs_info表
type HsInfo struct {
	gorm.Model
	Title    string `gorm:"unique;not null;comment:标题"`
	Url      string
	M3u8Url  string
	Platform string
	ClassId  int
	Page     int
	PhotoUrl string
	Location string
}

// 返回视频对象信息
type VideoInfo struct {
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
}

type User struct {
	gorm.Model
	Username string
	Password string
	Address  string
	Phone    string
}

type Title struct {
	title string
}

// 实现一个接口 重载两个方法
type Func interface {
	MaodouReq(...interface{})
	RequestPageInfo(...interface{})
}

// 旧方法实现
type Org struct {
}

// 新方法实现
type New struct {
}

func FindTitleList() {
	mysql, err := db.MysqlConfigure()
	if err != nil {
		fmt.Println("connent datebase err:", err)
	}
	titles := make([]Title, 3)
	mysql.Table("t_hs_info").Select([]string{"title"}).Scan(&titles)
	fmt.Println(titles)
}

// 同步redis数据 遍历redis数据
func Redis2Mysql() {
	keys := redis.GetKeyList()
	mysql, err := db.MysqlConfigure()
	if err != nil {
		fmt.Println("connent datebase err:", err)
	}
	for i, key := range keys {
		values, _ := redis.GetKey(key)
		hsInfo := HsInfo{}
		json.Unmarshal(utils.String2Bytes(values), &hsInfo)
		mysql.Create(&hsInfo)
		fmt.Printf("第%d个\n", i+1)
	}
}

// mysql数据同步redis
func Mysql2Redis() {
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

// redis查询包含数据打印
func PrintfCommon(page, num int, href, title string, row int64, platform string) {
	fmt.Printf("platform:[ %s ]-location:[ %d,%d ]-row:[ %d ]-href:[ %s ]-title:[ %s ]\n", platform, page, num+1, row, href, title)
}

// 请求接口传输数据
func RequestMysqlSave(hsInfo HsInfo) {
	url := "http://localhost:8520/sentinel_client_sale/hsInfo/save"
	method := "POST"

	json, _ := json.Marshal(hsInfo)
	fmt.Println(string(json))

	payload := strings.NewReader(strings.ToUpper(string(json)))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// 视频展示页面转为播放页面
func Display2Video(url_org, url_new string) string {
	html := ".html"
	index := strings.LastIndex(url_new, "/")
	s1 := url_new[index:]
	s2 := strings.Replace(s1, html, "", -1)
	return url_org + "/index.php/vod/play/id" + s2 + "/sid/1/nid/1" + html
}

// 根据页面转换请求url
func ConvertUrl(url string, page int) string {
	html := ".html"
	if page == 1 {
		return url
	}
	index := strings.Index(url, html)
	return url[:index] + "/page/" + strconv.Itoa(page) + html
}

// 请求播放页面拿到m3u8url
func PlayVideoM3u8Info(url string, location int) string {

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android............ecko) Chrome/92.0.4515.105 HuaweiBrowser/12.0.4.300 Mobile Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	reader, _ := goquery.NewDocumentFromReader(res.Body)

	var m3u8_url string

	reader.Find("script").Each(func(i int, selection *goquery.Selection) {
		title := selection.Text()
		if i == location {
			s2 := strings.Replace(title, "\\/", "/", -1)
			index := strings.Index(s2, "=")
			s3 := s2[index+1:]
			var videoInfo VideoInfo
			json.Unmarshal([]byte(s3), &videoInfo)
			m3u8_url = videoInfo.Url
		}
	})

	return m3u8_url

}
