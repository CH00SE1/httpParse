package open

import (
	"httpParse/db"
	"httpParse/guojiayibao"
	"time"
)

/**
 * @title xxx
 * @author xiongshao
 * @date 2022-07-13 09:28:40
 */

func T1001() {
	configure, _ := db.MysqlConfigure()
	for i := 1; i < 140; i++ {
		data := guojiayibao.GetData1001(i)
		for _, row := range data.Rows {
			if row.AuditAddTime.IsZero() {
				row.AuditAddTime = time.Now()
			}
			configure.Create(&row)
		}
	}
}

func T1002() {
	configure, _ := db.MysqlConfigure()
	for i := 1; i < 50; i++ {
		data1002 := guojiayibao.GetData1002(i)
		for _, data := range data1002.Rows {
			configure.Create(&data)
		}
	}
}

func T1003() {
	configure, _ := db.MysqlConfigure()
	for i := 1; i < 25; i++ {
		data := guojiayibao.GetData1003(i)
		for _, row := range data.Rows {
			configure.Create(&row)
		}
	}
}
