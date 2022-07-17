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
	for i := 1; i < 138; i++ {
		data := guojiayibao.GetData1001(i)
		for _, row := range data.Rows {
			if row.AuditAddTime.IsZero() {
				row.AuditAddTime = time.Now()
			}
			configure.Table("t_drug1001_tmp").Create(&row)
		}
	}
}

func T1002() {
	configure, _ := db.MysqlConfigure()
	for i := 1; i < 44; i++ {
		data1002 := guojiayibao.GetData1002(i)
		for _, data := range data1002.Rows {
			configure.Table("t_drug_info1002_tmp").Create(&data)
		}
	}
}

func T1003() {
	configure, _ := db.MysqlConfigure()
	for i := 3; i < 9; i++ {
		data := guojiayibao.GetData1003(i)
		for _, row := range data.Rows {
			configure.Table("t_good_info1003").Create(&row)
		}
	}
}
