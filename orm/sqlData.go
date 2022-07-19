package orm

import (
	"database/sql"
	"fmt"
	dbOracle "httpParse/db"
)

/**
 * @title orm结构
 * @author xiongshao
 * @date 2022-07-18 09:32:16
 */

func query(sqlStr string, db *sql.DB) (list []map[string]interface{}) {
	rows, _ := db.Query(sqlStr)
	//字段名称
	columns, _ := rows.Columns()
	//多少个字段
	length := len(columns)
	//每一行字段的值
	values := make([]sql.RawBytes, length)
	//保存的是values的内存地址
	pointer := make([]interface{}, length)
	//
	for i := 0; i < length; i++ {
		pointer[i] = &values[i]
	}
	//
	for rows.Next() {
		//把参数展开，把每一行的值存到指定的内存地址去，循环覆盖，values也就跟着被赋值了
		err := rows.Scan(pointer...)
		if err != nil {
			fmt.Println(err)
			return
		}
		//每一行
		row := make(map[string]interface{})
		for i := 0; i < length; i++ {
			row[columns[i]] = string(values[i])
		}
		list = append(list, row)
	}
	_ = rows.Close()
	return
}

func GetOracleData(sqlStr, oracleName string) (list []map[string]interface{}) {
	if oracleName == dbOracle.ERP {
		db, _ := dbOracle.ErpOracleConfig()
		return query(sqlStr, db)
	}
	if oracleName == dbOracle.DS {
		db, _ := dbOracle.DsOracleConfig()
		return query(sqlStr, db)
	}
	return nil
}
