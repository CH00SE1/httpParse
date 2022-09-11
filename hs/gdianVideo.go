package hs

import (
	"encoding/json"
	"fmt"
	"httpParse/redis"
	"httpParse/utils"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

/**
 * @title G.视频
 * @author xiongshao
 * @date 2022-07-13 11:38:26
 */

// 视频主页
const g_url = "https://www.gdian125.com"

type PhotoDao struct {
	Msg     string `json:"msg"`
	Code    string `json:"code"`
	IsLogin int    `json:"is_login"`
	Data    struct {
		Data []struct {
			MovieId         int           `json:"movie_id"`
			MovieName       string        `json:"movie_name"`
			MovieCover      string        `json:"movie_cover"`
			MovieLong       string        `json:"movie_long"`
			WatchCount      int           `json:"watch_count"`
			ForeverPoint    int           `json:"forever_point"`
			NeedLogin       int           `json:"need_login"`
			NeedVip         int           `json:"need_vip"`
			MovieScore      int           `json:"movie_score"`
			MoviePreviewImg interface{}   `json:"movie_preview_img"`
			PushTime        int           `json:"push_time"`
			JavNumber       interface{}   `json:"jav_number"`
			ActorIds        []interface{} `json:"actor_ids"`
			ListType        int           `json:"list_type"`
			Labels          []string      `json:"labels"`
		} `json:"data"`
		Total int `json:"total"`
	} `json:"data"`
}

type GM3u8VideoRes struct {
	Msg     string `json:"msg"`
	Code    string `json:"code"`
	IsLogin int    `json:"is_login"`
	Data    struct {
		MovieId         int           `json:"movie_id"`
		MovieName       string        `json:"movie_name"`
		MovieCover      string        `json:"movie_cover"`
		MovieLong       string        `json:"movie_long"`
		WatchCount      int           `json:"watch_count"`
		ForeverPoint    int           `json:"forever_point"`
		NeedLogin       int           `json:"need_login"`
		NeedVip         int           `json:"need_vip"`
		MovieScore      int           `json:"movie_score"`
		MoviePreviewImg interface{}   `json:"movie_preview_img"`
		PushTime        int           `json:"push_time"`
		JavNumber       interface{}   `json:"jav_number"`
		ActorIds        []interface{} `json:"actor_ids"`
		MovieSecond     int           `json:"movie_second"`
		Introduction    string        `json:"introduction"`
		OnedayPoint     int           `json:"oneday_point"`
		LikeCount       int           `json:"like_count"`
		CollectCount    int           `json:"collect_count"`
		CommentCount    int           `json:"comment_count"`
		MovieM3U8Url    []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"movie_m3u8_url"`
		Labels           []string `json:"labels"`
		IsWatch          int      `json:"is_watch"`
		CanWatchTime     int      `json:"can_watch_time"`
		Discount         int      `json:"discount"`
		MovieMp4         string   `json:"movie_mp4"`
		IsLike           int      `json:"is_like"`
		IsCollect        int      `json:"is_collect"`
		HistoryLong      int      `json:"history_long"`
		VipLineNum       int      `json:"vip_line_num"`
		HistoryLine      int      `json:"history_line"`
		UserCouponsState int      `json:"user_coupons_state"`
		SuggestionList   []struct {
			MovieId         int           `json:"movie_id"`
			MovieName       string        `json:"movie_name"`
			MovieCover      string        `json:"movie_cover"`
			MovieLong       string        `json:"movie_long"`
			WatchCount      int           `json:"watch_count"`
			ForeverPoint    int           `json:"forever_point"`
			NeedLogin       int           `json:"need_login"`
			NeedVip         int           `json:"need_vip"`
			MovieScore      int           `json:"movie_score"`
			MoviePreviewImg interface{}   `json:"movie_preview_img"`
			PushTime        int           `json:"push_time"`
			JavNumber       interface{}   `json:"jav_number"`
			ActorIds        []interface{} `json:"actor_ids"`
			ListType        int           `json:"list_type"`
			Labels          []string      `json:"labels"`
		} `json:"suggestion_list"`
	} `json:"data"`
}

// G.请求url拿到数据
func GRequest(page int) {
	url := g_url + "/apiv2/video/search?is_av=0&sort=latest&page=" + strconv.Itoa(page) + "&num=20"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "www.gdian169.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("agent", "PC")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cookie", "PHPSESSID=f811ef900e894b1592a3d9beba2912bf; AWSELB=7DA38D4B14F1329125C38FB72B7398A7663A9656C0B80E962609C0ECF41DA7039E9B85B63FB124B19867940202770B4197ED805C4A114768600E593194FD7566CBB0C62C67; Edge-Sticky=jESAbOBzx2V6Bn9vuffmAw==; Hm_lvt_b23c8eb27081cd4e5308c6f1df7643cb=1657683383; _gid=GA1.2.980495506.1657683383; XLA_CI=6038dab041ca67e5ee2940eca0234bf4; _ga=GA1.1.1377953948.1657683383; AWSALB=jWwIpnawSFgh1Ta9aH/e5aOpqZR+9vG3GmncwTRc5taMnQpw5HB9GyFsVhifpDstWB8IAt8XWpRP7G2B+J0IbNfKW6FK+J8R9sv9sO2KvzzMZ8SrbbQOIsekdqF2; fish_session=Qz3fditrY5HQpZYCOOuEL9haO7wIp6ol27Pu7UrM; Hm_lpvt_b23c8eb27081cd4e5308c6f1df7643cb=1657683967; _ga_TM83R4D3QS=GS1.1.1657683395.1.1.1657683967.0; fish_session=Qz3fditrY5HQpZYCOOuEL9haO7wIp6ol27Pu7UrM")
	req.Header.Add("referer", "https://www.gdian169.com/pc/video")
	req.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.49")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var photoDao PhotoDao
	json.Unmarshal(body, &photoDao)
	//mysql, _ := db.MysqlConfigure()
	var hsInfos []*HsInfo
	for num, datum := range photoDao.Data.Data {
		// 数据保存
		hsInfo := resultObejectInfo(datum.MovieName, datum.MovieId, page)
		if hsInfo != nil {
			info := hsInfo.(HsInfo)
			hsInfos = append(hsInfos, &info)
		}
		// 图片下载
		savePhotoInfo(datum.MovieCover, datum.MovieName, num+1)
	}
}

// 获取video信息
func m3u8VideoInfo(movie_id int) (string, GM3u8VideoRes) {
	url := g_url + "/apiv2/video/" + strconv.Itoa(movie_id)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "www.gdian169.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("agent", "PC")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cookie", "PHPSESSID=f811ef900e894b1592a3d9beba2912bf; AWSELB=7DA38D4B14F1329125C38FB72B7398A7663A9656C0B80E962609C0ECF41DA7039E9B85B63FB124B19867940202770B4197ED805C4A114768600E593194FD7566CBB0C62C67; Edge-Sticky=jESAbOBzx2V6Bn9vuffmAw==; Hm_lvt_b23c8eb27081cd4e5308c6f1df7643cb=1657683383; _gid=GA1.2.980495506.1657683383; XLA_CI=6038dab041ca67e5ee2940eca0234bf4; _ga=GA1.1.1377953948.1657683383; fish_session=Qz3fditrY5HQpZYCOOuEL9haO7wIp6ol27Pu7UrM; AWSALB=toMWmnTbFXdK+Gg1icKfHQqy5EiA9cSH2TA2nf63M4jdhzF6RHOJOntiH5qnDlH8vYtaa1ntu6GPDzn+tC8peZcd+cCPzcVNMldCwsCd9fXgFuDqCZKuQYIalZ2w; Hm_lpvt_b23c8eb27081cd4e5308c6f1df7643cb=1657687593; _ga_TM83R4D3QS=GS1.1.1657690222.3.1.1657690230.0; fish_session=Qz3fditrY5HQpZYCOOuEL9haO7wIp6ol27Pu7UrM")
	req.Header.Add("referer", "https://www.gdian169.com/pc/video")
	req.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.49")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var gM3u8VideoRes GM3u8VideoRes
	json.Unmarshal(body, &gM3u8VideoRes)
	return url, gM3u8VideoRes
}

// 下载视频的图片
func savePhotoInfo(url, fileName string, num int) {
	method := "GET"

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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes2String := utils.Bytes2String(body)
	utils.CreateFile(&bytes2String, "C:\\Users\\Administrator\\Desktop\\photo_G\\",
		"G"+time.Now().Format("[2006-01-02-15-04-05]-")+fileName+strconv.Itoa(num), ".jpg")
}

// 数据保存mysql
func resultObejectInfo(MovieName string, MovieId, page int) interface{} {
	row := redis.KeyExists(MovieName)
	if row != 1 {
		watchUrl, videoInfo := m3u8VideoInfo(MovieId)
		var m3u8_url string
		for _, urlDao := range videoInfo.Data.MovieM3U8Url {
			if urlDao.Name == "普通线路" {
				m3u8_url = urlDao.Url
			}
		}
		var labels string
		for _, label := range videoInfo.Data.Labels {
			labels += ("`" + label + "`")
		}
		labels += videoInfo.Data.MovieLong
		hsInfo := HsInfo{Title: videoInfo.Data.MovieName,
			Url:      watchUrl,
			M3u8Url:  g_url + m3u8_url,
			ClassId:  page,
			Platform: "G.视频[" + labels + "]",
			Page:     page,
			Location: videoInfo.Data.MovieCover}
		marshal, _ := json.Marshal(hsInfo)
		redis.SetKey(videoInfo.Data.MovieName, marshal)
		fmt.Println("set key title:", hsInfo.Title)
		return hsInfo
	}
	return nil
}
