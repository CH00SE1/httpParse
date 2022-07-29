package hs

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"httpParse/db"
	"httpParse/redis"
	"httpParse/utils"
	"io/ioutil"
	"net/http"
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
	Location string
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

// redis查询包含数据打印
func PrintfCommon(page, num int, href, title string, row int64, platform string) {
	fmt.Printf("\nplatform:[%s]-location:[%d,%d]-row:[%d]\nhref:{%s}\ntitle:{%s}\n", platform, page, num, row, href, title)
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
