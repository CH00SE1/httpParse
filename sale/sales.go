package sale

import (
	"fmt"
	DB "httpParse/db"
	"httpParse/orm"
	"strconv"
)

/**
 * @title sales销售查询
 * @author xiongshao
 * @date 2022-07-18 15:54:19
 */

type SaleInfoRes struct {
	Title            string            `json:"title"`
	ShopId           int               `json:"shop_id"`
	Type             string            `json:"type"`
	SalesInfoDetails []SalesInfoDetail `json:"sales_info_details"`
	SalesName        []string          `json:"sales_name"`
	Classification   []string          `json:"classification"`
}

type SalesInfoDetail struct {
	Name       string               `json:"name"`
	SumMoney   float64              `json:"sum_money"`
	SumFlMoney float64              `json:"sum_fl_money"`
	Fl         []map[string]float64 `json:"fl"`
}

func SaleFind(shopId int, Type string) SaleInfoRes {
	datas, type_status, dataTime := getData(shopId, Type)
	names := getMap("EMPLOYEENAME", datas)
	fls := getMap("FL", datas)
	var salesInfoDetails []SalesInfoDetail
	// 二层
	for _, name := range names {
		n, sale := findNameAllSale(name, datas)
		var allMl float64
		var list []map[string]float64
		for _, fl := range fls {
			classMl, aMl := findClassMl(name, fl, datas)
			allMl += aMl
			list = append(list, classMl)
		}
		detail := SalesInfoDetail{
			Name:       n,
			SumMoney:   sale,
			SumFlMoney: allMl,
			Fl:         list,
		}
		salesInfoDetails = append(salesInfoDetails, detail)
	}
	// 一层
	var saleInfoRes = SaleInfoRes{
		ShopId:           shopId,
		Title:            dataTime,
		Type:             type_status,
		SalesName:        names,
		Classification:   fls,
		SalesInfoDetails: salesInfoDetails,
	}
	return saleInfoRes
}

// 根据人员毛利统计销售
func findClassMl(name, fl string, lists []map[string]interface{}) (map[string]float64, float64) {
	maps := make(map[string]float64)
	var classFl float64
	for _, list := range lists {
		if name == list["EMPLOYEENAME"] && fl == list["FL"] && list["CLASSNAME"] == "" {
			str := list["REALMONEY"].(string)
			float, _ := strconv.ParseFloat(str, 64)
			classFl += ml(fl, float)
		}
		maps[fl] = classFl
	}
	return maps, classFl
}

// 查询人员的总销售
func findNameAllSale(name string, lists []map[string]interface{}) (string, float64) {
	var allSale float64
	for _, list := range lists {
		if name == list["EMPLOYEENAME"] {
			str := list["REALMONEY"].(string)
			float, _ := strconv.ParseFloat(str, 64)
			allSale += float
		}
	}
	return name, allSale
}

// 获取数据来源
func getData(shopId int, dateTime string) ([]map[string]interface{}, string, string) {
	var Type string
	// fl is not null and realmoney is not null and
	data := "select realmoney,employeename,fl,classname from gresa_sa_detail_query_v " +
		"where PLACEPOINTID = " + strconv.Itoa(shopId)
	switch dateTime {
	case "day":
		data += " and trunc(sysdate) = trunc(CREDATE) "
		Type = "当天"
	case "yesterday":
		data += " and TO_CHAR(CREDATE,'YYYY-MM-DD') = TO_CHAR(SYSDATE-1,'YYYY-MM-DD') "
		Type = "昨天"
	case "month":
		data += " and CREDATE >= trunc(sysdate,'MM') and CREDATE <= last_day(sysdate) "
		Type = "当月"
	case "lastmonth":
		data += " and TO_CHAR(CREDATE,'YYYY-MM') = TO_CHAR(ADD_MONTHS(SYSDATE,-1),'YYYY-MM') "
		Type = "上月"
	case "year":
		data += " and CREDATE >= trunc(sysdate,'YYYY') and CREDATE <= ADD_MONTHS(trunc(sysdate,'YYYY'),12) - 1 "
		Type = "当年 "
	}
	fmt.Println("SQL:", data)
	return orm.GetOracleData(data, DB.ERP), Type, dateTime
}

// 根据key获取
func getMap(key string, list []map[string]interface{}) []string {
	maps := make(map[interface{}]interface{})
	for _, datum := range list {
		if datum[key] != "" {
			NAME := datum[key]
			maps[NAME] = nil
		}
	}
	return getArray(maps)
}

// map key标签转为[]string
func getArray(m map[interface{}]interface{}) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	keys := make([]string, 0, len(m))
	for k := range m {
		str := fmt.Sprintf("%v", k)
		keys = append(keys, str)
	}
	return keys
}

// 计算分类毛利
func ml(fl string, num float64) float64 {
	var n float64
	switch fl {
	case "A":
		n = num * 0.1
	case "B":
		n = num * 0.08
	case "C":
		n = num * 0.05
	case "D":
		n = num * 0.03
	case "E":
		n = num * 0
	}
	return n
}
