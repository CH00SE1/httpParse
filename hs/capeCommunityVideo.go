package hs

import (
	"encoding/json"
	"fmt"
	"httpParse/db"
	"httpParse/redis"
	"io"
	"net/http"
	"strconv"
)

/**
 * @title 海角社区
 * @author xiongshao
 * @date 2022-07-16 11:23:22
 */

// 主题详情
type CapePageInfo struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Success   bool   `json:"success"`
	Data      struct {
		Page struct {
			Page  int `json:"page"`
			Limit int `json:"limit"`
			Total int `json:"total"`
		} `json:"page"`
		Results []struct {
			TopicId int `json:"topicId"`
			User    struct {
				Id             int    `json:"id"`
				Nickname       string `json:"nickname"`
				Avatar         string `json:"avatar"`
				Vip            int    `json:"vip"`
				Certified      bool   `json:"certified"`
				CertVideo      bool   `json:"certVideo"`
				CertProfessor  bool   `json:"certProfessor"`
				Famous         bool   `json:"famous"`
				DiamondConsume int    `json:"diamondConsume"`
				Title          struct {
					Id         int    `json:"id"`
					Name       string `json:"name"`
					Consume    int    `json:"consume"`
					ConsumeEnd int    `json:"consumeEnd"`
					Icon       string `json:"icon"`
				} `json:"title"`
			} `json:"user"`
			Node struct {
				NodeId      int    `json:"nodeId"`
				ParentId    int    `json:"parentId"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				SortNo      int    `json:"sortNo"`
				Description string `json:"description"`
				Display     int    `json:"display"`
				VipLimit    int    `json:"vipLimit"`
				ExternalUrl string `json:"external_url"`
			} `json:"node"`
			Tags []struct {
				TagId   int    `json:"tagId"`
				TagName string `json:"tagName"`
			} `json:"tags"`
			Title           string `json:"title"`
			Type            int    `json:"type"`
			LiteContent     string `json:"liteContent"`
			LastCommentTime string `json:"lastCommentTime"`
			ViewCount       int    `json:"viewCount"`
			CommentCount    int    `json:"commentCount"`
			LikeCount       int    `json:"likeCount"`
			Status          int    `json:"status"`
			CreateTime      string `json:"createTime"`
			Attachments     []struct {
				Id        int    `json:"id"`
				RemoteUrl string `json:"remoteUrl"`
				Category  string `json:"category"`
				Status    int    `json:"status"`
				CoverUrl  string `json:"coverUrl,omitempty"`
			} `json:"attachments"`
			IsCream        bool   `json:"is_cream"`
			IsTop          bool   `json:"is_top"`
			IsSubscribe    bool   `json:"is_subscribe"`
			IsHot          bool   `json:"is_hot"`
			IsOriginal     bool   `json:"is_original"`
			Remarks        string `json:"remarks"`
			TopExpiresTime string `json:"topExpiresTime"`
		} `json:"results"`
	} `json:"data"`
}

// 页面详情
type CapePageByIdInfo struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Success   bool   `json:"success"`
	Data      struct {
		TopicId int `json:"topicId"`
		User    struct {
			Id             int    `json:"id"`
			Nickname       string `json:"nickname"`
			Avatar         string `json:"avatar"`
			Vip            int    `json:"vip"`
			Certified      bool   `json:"certified"`
			CertVideo      bool   `json:"certVideo"`
			CertProfessor  bool   `json:"certProfessor"`
			Famous         bool   `json:"famous"`
			DiamondConsume int    `json:"diamondConsume"`
			Title          struct {
				Id         int    `json:"id"`
				Name       string `json:"name"`
				Consume    int    `json:"consume"`
				ConsumeEnd int    `json:"consumeEnd"`
				Icon       string `json:"icon"`
			} `json:"title"`
		} `json:"user"`
		Node struct {
			NodeId      int    `json:"nodeId"`
			ParentId    int    `json:"parentId"`
			Name        string `json:"name"`
			Icon        string `json:"icon"`
			SortNo      int    `json:"sortNo"`
			Description string `json:"description"`
			Display     int    `json:"display"`
			VipLimit    int    `json:"vipLimit"`
			ExternalUrl string `json:"external_url"`
		} `json:"node"`
		Tags []struct {
			TagId   int    `json:"tagId"`
			TagName string `json:"tagName"`
		} `json:"tags"`
		Title           string `json:"title"`
		Type            int    `json:"type"`
		LiteContent     string `json:"liteContent"`
		LastCommentTime string `json:"lastCommentTime"`
		ViewCount       int    `json:"viewCount"`
		CommentCount    int    `json:"commentCount"`
		LikeCount       int    `json:"likeCount"`
		Status          int    `json:"status"`
		CreateTime      string `json:"createTime"`
		Attachments     []struct {
			Id              int    `json:"id"`
			RemoteUrl       string `json:"remoteUrl"`
			Category        string `json:"category"`
			Status          int    `json:"status"`
			CoverUrl        string `json:"coverUrl,omitempty"`
			VideoTimeLength int    `json:"video_time_length,omitempty"`
		} `json:"attachments"`
		IsCream        bool        `json:"is_cream"`
		IsTop          bool        `json:"is_top"`
		IsSubscribe    bool        `json:"is_subscribe"`
		IsHot          bool        `json:"is_hot"`
		IsOriginal     bool        `json:"is_original"`
		Remarks        string      `json:"remarks"`
		TopExpiresTime string      `json:"topExpiresTime"`
		Content        string      `json:"content"`
		Reward         interface{} `json:"reward"`
		Sale           struct {
			MoneyType     int  `json:"money_type"`
			Amount        int  `json:"amount"`
			IsBuy         bool `json:"is_buy"`
			BuyIndex      int  `json:"buy_index"`
			BuyCount      int  `json:"buyCount"`
			BuyCountFloor int  `json:"buyCountFloor"`
		} `json:"sale"`
		Doors interface{} `json:"doors"`
	} `json:"data"`
}

// 主页 url
const (
	cape_url           = "https://hjf2d1.com/api/topic/"
	hot, news          = "hot/topics?", "node/news?"
	dashiji, yuanchuan = "node/topics?type=1&nodeId=258&", "node/topics?type=7&"
	jinghua, new       = "node/topics?type=3&nodeId=0&", "node/topics?type=1&nodeId=0&"
)

func (org Org) RequestPageInfo(page int) {
	url := cape_url + dashiji + "page=" + strconv.Itoa(page)
	method := "GET"

	fmt.Printf("\nurl : %s\n", url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "hjf2d1.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cookie", "_ga=GA1.1.26407275.1657940574; NOTLOGIN=NOTLOGIN; _ga_H4G4E5X3FL=GS1.1.1657940573.1.1.1657942139.0")
	req.Header.Add("pcver", "220708143229")
	req.Header.Add("referer", "https://hjf2d1.com/")
	req.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var capePageInfo CapePageInfo
	json.Unmarshal(body, &capePageInfo)
	for _, result := range capePageInfo.Data.Results {
		requestPageByIdInfo(result.TopicId, page)
	}
}

func (new New) RequestPageInfo(page int, status string) {
	url := cape_url + status + "page=" + strconv.Itoa(page)
	method := "GET"

	fmt.Printf("\nurl : %s\n", url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "hjf2d1.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cookie", "_ga=GA1.1.26407275.1657940574; NOTLOGIN=NOTLOGIN; _ga_H4G4E5X3FL=GS1.1.1657940573.1.1.1657942139.0")
	req.Header.Add("pcver", "220708143229")
	req.Header.Add("referer", "https://hjf2d1.com/")
	req.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var capePageInfo CapePageInfo
	json.Unmarshal(body, &capePageInfo)
	for _, result := range capePageInfo.Data.Results {
		requestPageByIdInfo(result.TopicId, page)
	}
}

func requestPageByIdInfo(page_id, page int) {
	url := cape_url + strconv.Itoa(page_id)
	method := "GET"

	fmt.Printf("\nurl : %s\n", url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "hjf2d1.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cookie", "_ga=GA1.1.26407275.1657940574; NOTLOGIN=NOTLOGIN; _ga_H4G4E5X3FL=GS1.1.1657940573.1.1.1657941396.0")
	req.Header.Add("pcver", "220708143229")
	req.Header.Add("referer", "https://hjf2d1.com/post/details?pid=370111")
	req.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var capePageByIdInfo CapePageByIdInfo
	json.Unmarshal(body, &capePageByIdInfo)
	dataSave(capePageByIdInfo, url, page)
}

// 数据保存
func dataSave(capePageByIdInfo CapePageByIdInfo, url string, page int) {
	title := capePageByIdInfo.Data.Title
	var m3u8_url string
	for _, attachment := range capePageByIdInfo.Data.Attachments {
		if attachment.Category == "video" {
			m3u8_url = attachment.RemoteUrl
		}
	}
	if len(m3u8_url) != 0 {
		db, _ := db.MysqlConfigure()
		row := redis.KeyExists(title)
		if row != 1 {
			hsInfo := HsInfo{
				Title:    title,
				Url:      url,
				M3u8Url:  m3u8_url,
				ClassId:  capePageByIdInfo.Data.Type,
				Platform: "海角社区",
				Page:     page,
				Location: "视频上线时间:" + capePageByIdInfo.Data.CreateTime,
			}
			marshal, err := json.Marshal(hsInfo)
			if err != nil {
				fmt.Println("hsInfo json 序列化失败")
			}
			redis.SetKey(title, marshal)
			db.Create(&hsInfo).Callback()
		} else {
			fmt.Printf("page:%d title:%s\n", page, title)
			PrintfCommon(page, 1, url, title, row, "海角社区")
		}
	} else {
		fmt.Printf("*****************%s*****************\n", title)
		/*content := capePageByIdInfo.Data.Content
		dom, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(content)))
		if err != nil {
			log.Fatal(err)
		}
		var Text string
		num := 1
		dom.Find("p").Each(func(i int, selection *goquery.Selection) {
			text := utils.StringStrip(selection.Text())
			if len(text) != 0 {
				Text += strconv.Itoa(num) + "." + text + "\n"
				num++
			}
		})
		if num != 1 {
			utils.CreateFile(&Text, "D:\\海角社区\\", title, ".txt")
		}*/
	}
}
