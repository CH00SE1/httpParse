package main

import (
	"httpParse/db"
	"httpParse/guojiayibao"
	"httpParse/hs"
	"sync"
	"time"
)

func T1001() {
	for i := 1; i < 22; i++ {
		data := guojiayibao.GetData1001(i)
		configure, _ := db.MysqlConfigure()
		for _, row := range data.Rows {
			if row.AuditAddTime.IsZero() {
				row.AuditAddTime = time.Now()
			}
			configure.Table("t_drug1001_tmp").Create(&row)
		}
	}
}

func T1002() {
	for i := 1; i < 44; i++ {
		data1002 := guojiayibao.GetData1002(i)
		configure, _ := db.MysqlConfigure()
		for _, data := range data1002.Rows {
			configure.Table("t_drug_info1002_tmp").Create(&data)
		}
	}
}

func T1003() {
	for i := 1; i < 30; i++ {
		data := guojiayibao.GetData1003(i)
		configure, _ := db.MysqlConfigure()
		for _, row := range data.Rows {
			configure.Table("t_good_info1003").Create(&row)
		}
	}
}

func Hs() {
	for i := 2; i < 10; i++ {
		hs.ExampleScrape(28, i)
	}
}

// 全局变量
var wg sync.WaitGroup
var tag = 1

func Tpaoyou0() {
	defer wg.Done()
	for i := 200; i < 230; i++ {
		hs.Paoyou(tag, i)
	}
}

func Tpaoyou1() {
	defer wg.Done()
	for i := 230; i < 260; i++ {
		hs.Paoyou(tag, i)
	}
}

func Tpaoyou2() {
	defer wg.Done()
	for i := 260; i < 300; i++ {
		hs.Paoyou(tag, i)
	}
}

func syncTpaoyou() {
	wg.Add(1)
	go Tpaoyou0()
	go Tpaoyou1()
	go Tpaoyou2()
	wg.Wait()
}

func main() {
	//gin := gin.Default()
	//gin.GET("/getData/:page/:start", src.SaveInfo)
	//gin.Run(":8500")
	//db.AutoCreateTable(xml.XmlInfo{})
	//T1002()
	syncTpaoyou()
}
