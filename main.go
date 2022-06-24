package main

import (
	"fmt"
	"httpParse/li5apuu7"
)

func main() {
	//gin := gin.Default()
	//gin.GET("/getData/:page/:start", src.SaveInfo)
	//gin.Run(":8500")
	//db.AutoCreateTable(guojiayibao.Drug1001Tmp{})
	for i := 114; i < 500; i++ {
		_, i2 := li5apuu7.ExampleScrape(21, i)
		fmt.Println("第 -------------- ", i2, " -------------- 页")
	}
	// 1001
	//for i := 1; i < 22; i++ {
	//	data := guojiayibao.GetData1001(i)
	//	configure, _ := db.MysqlConfigure()
	//	for _, row := range data.Rows {
	//		if row.AuditAddTime.IsZero() {
	//			row.AuditAddTime = time.Now()
	//		}
	//		configure.Table("t_drug1001_tmp").Create(&row)
	//	}
	//}
	// 1003
	//for i := 1; i < 30; i++ {
	//	data := guojiayibao.GetData1003(i)
	//	configure, _ := db.MysqlConfigure()
	//	for _, row := range data.Rows {
	//		configure.Table("t_good_info1003").Create(&row)
	//	}
	//}
}
