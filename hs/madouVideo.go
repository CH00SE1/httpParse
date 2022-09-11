package hs

import (
	"encoding/json"
	"fmt"
	"httpParse/db"
	"httpParse/redis"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
 * @title madou视频
 * @author xiongshao
 * @date 2022-07-14 15:40:28
 */

const (
	Tv91_url, maodou_url, quanyuan_url = "https://nnp35.com/upload_json_live", "https://jsonmdtv.md29.tv/upload_json_live", "https://json.wtjfjil.cn/upload_json_live"
)

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

// 区分平台
func platform(Type string) string {
	var platfrom string
	switch Type {
	case Tv91_url:
		platfrom = "91视频"
	case maodou_url:
		platfrom = "麻豆视频"
	case quanyuan_url:
		platfrom = "A1LT1N"
	}
	return platfrom
}

// 根据平台转发请求xxx.json
func convertThreeUrl(urlType string, page int) (string, string) {

	date := strings.Replace(time.Now().Format("2006-01-02"), "-", "", -1)

	var hour int
	if time.Now().Hour()%2 == 1 {
		hour = time.Now().Hour() - 1
	} else {
		hour = time.Now().Hour()
	}

	str_hour := ""
	if hour < 10 {
		str_hour += "0" + strconv.Itoa(hour)
	} else {
		str_hour += strconv.Itoa(hour)
	}

	var url string
	switch urlType {
	case Tv91_url:
		url = Tv91_url + "/" + date + "/videolist_" + date + "_" + str_hour + "_2_-_-_100_" + strconv.Itoa(page) + ".json"
	case maodou_url:
		url = maodou_url + "/" + date + "/videolist_" + date + "_" + str_hour + "_2_-_-_100_" + strconv.Itoa(page) + ".json"
	case quanyuan_url:
		url = quanyuan_url + "/" + date + "/videolist_zh-cn_" + date + "_" + str_hour + "_-_-_-_50_" + strconv.Itoa(page) + ".json"
	}

	return url, urlType
}

// ------------------------------------------------ madou ------------------------------------------------
// 初始方法
func (org Org) MaodouReq(page int) ([]byte, string, string) {

	urlType := Tv91_url

	url, Type := convertThreeUrl(urlType, page)

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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected EOF") && len(body) != 0 {
			log.Fatal(err, page)
		}
	}
	return body, Type, urlType
}

// 新方法
func (new New) MaodouReq(page int, platform string) ([]byte, string, string) {

	url, Type := convertThreeUrl(platform, page)

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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected EOF") && len(body) != 0 {
			log.Fatal(err, page)
		}
	}
	return body, Type, platform
}

// 数据转化存储
func DataParseSave(body []byte, Type, urlType string) {
	var maDouDao MaDouDao
	json.Unmarshal(body, &maDouDao)

	datas := maDouDao.Data
	db, _ := db.MysqlConfigure()
	for i, data := range datas {
		row := redis.KeyExists(data.Title)
		if row != 1 {
			hsInfo := HsInfo{
				Title:    data.Title,
				Url:      urlType + data.TestVideoUrl,
				M3u8Url:  strings.Replace(data.VideoUrl, "\\/", "/", -1),
				ClassId:  maDouDao.CurrentPage,
				Platform: platform(Type) + "*" + data.ComefromTitle,
				Page:     maDouDao.CurrentPage,
				Location: "[" + strconv.Itoa((i+1)/6+1) + "," + strconv.Itoa((i+1)%6+1) + "]",
				PhotoUrl: data.Panorama,
			}
			marshal, err := json.Marshal(hsInfo)
			if err != nil {
				fmt.Println("hsInfo json 序列化失败")
			}
			redis.SetKey(data.Title, marshal)
			db.Create(&hsInfo).Callback()
		} else {
			PrintfCommon(maDouDao.CurrentPage, i+1, urlType+data.TestVideoUrl, data.Title, row, platform(Type))
		}
	}
	// 创建文件
	//bytes2String := utils.Bytes2String(body)
	//utils.CreateFile(&bytes2String, "D:\\MadouData\\response\\", "madou_"+
	//time.Now().Format("[2006-01-02-15-04-05]_page_")+strconv.Itoa(maDouDao.PerPage), ".json")
}
