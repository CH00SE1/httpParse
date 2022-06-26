package main

import (
	"fmt"
	"httpParse/db"
	"httpParse/guojiayibao"
	"httpParse/yellow"
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

func Tpaoyou() {
	for i := 1; i < 100; i++ {
		yellow.Paoyou(1, i)
	}

}

func TdataJid() {
	yellow.GetDataJid("https://paoyou.ml/play/422100.html")
}

func TGetM3U8URl() {
	rl := yellow.GetM3U8URl("425326")
	fmt.Println(rl)
}

func main() {
	//gin := gin.Default()
	//gin.GET("/getData/:page/:start", src.SaveInfo)
	//gin.Run(":8500")
	//db.AutoCreateTable(guojiayibao.DrugInfo1002Tmp{})
	//T1002()
	//TGetM3U8URl()
	Tpaoyou()
}
