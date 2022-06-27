package main

import (
	"httpParse/db"
	"httpParse/guojiayibao"
	"httpParse/yellow"
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
		yellow.ExampleScrape(28, i)
	}
}

var wg sync.WaitGroup

func Tpaoyou0() {
	defer wg.Done()
	for i := 1; i < 30; i++ {
		yellow.Paoyou(3, i)
	}
}

func Tpaoyou1() {
	defer wg.Done()
	for i := 30; i < 60; i++ {
		yellow.Paoyou(3, i)
	}
}

func Tpaoyou2() {
	defer wg.Done()
	for i := 60; i < 101; i++ {
		yellow.Paoyou(3, i)
	}
}

func SyncTpaoyou() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Tpaoyou0()
		go Tpaoyou1()
		go Tpaoyou2()
	}
	wg.Wait()
}

func main() {
	//gin := gin.Default()
	//gin.GET("/getData/:page/:start", src.SaveInfo)
	//gin.Run(":8500")
	//db.AutoCreateTable(xml.XmlInfo{})
	//T1002()
	//TGetM3U8URl()
	SyncTpaoyou()
}
