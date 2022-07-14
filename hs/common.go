package hs

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"httpParse/db"
	"httpParse/redis"
	"httpParse/utils"
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
