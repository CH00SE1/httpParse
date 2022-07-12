package guojiayibao

import (
	"encoding/json"
	"fmt"
	"httpParse/db"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
 * @title 1104
 * @author xiongshao
 * @date 2022-07-12 11:21:03
 */

type Doc1104 struct {
	Records int `json:"records"`
	Total   int `json:"total"`
	Rows    []struct {
		SpecificationId    interface{} `json:"specificationId"`
		SpecificationCode  string      `json:"specificationCode"`
		Catalogcode        string      `json:"catalogcode"`
		Catalogname1       string      `json:"catalogname1"`
		Catalogname2       string      `json:"catalogname2"`
		Catalogname3       string      `json:"catalogname3"`
		Commonnamecode     string      `json:"commonnamecode"`
		Commonname         string      `json:"commonname"`
		Matrialcode        string      `json:"matrialcode"`
		Matrial            string      `json:"matrial"`
		Characteristiccode string      `json:"characteristiccode"`
		Characteristic     string      `json:"characteristic"`
		SeparateCharges    interface{} `json:"separateCharges"`
		PaymentType        interface{} `json:"paymentType"`
		PaymentStandard    interface{} `json:"paymentStandard"`
		Isusing            int         `json:"isusing"`
		SpecificationType  interface{} `json:"specificationType"`
		ProductStatusS     interface{} `json:"productStatusS"`
		ProductCount       interface{} `json:"productCount"`
		CompCount          interface{} `json:"compCount"`
		RegCount           interface{} `json:"regCount"`
		GoodsCount         interface{} `json:"goodsCount"`
	} `json:"rows"`
	Page        int         `json:"page"`
	Count       int         `json:"count"`
	FirstResult int         `json:"firstResult"`
	MaxResults  int         `json:"maxResults"`
	Success     bool        `json:"success"`
	Result      interface{} `json:"result"`
	Conditions  struct {
		OrderColumn int `json:"orderColumn"`
	} `json:"conditions"`
	Msg       interface{} `json:"msg"`
	Form      interface{} `json:"form"`
	Code      int         `json:"code"`
	OperCount int         `json:"operCount"`
	Sord      string      `json:"sord"`
	Sidx      string      `json:"sidx"`
	Orderby   interface{} `json:"orderby"`
}

type Dao1104 struct {
	Records int `json:"records"`
	Total   int `json:"total"`
	Rows    []struct {
		ReleaseVersion     int         `json:"releaseVersion"`
		CompanyName        interface{} `json:"companyName"`
		Catalogname1       interface{} `json:"catalogname1"`
		Catalogname2       interface{} `json:"catalogname2"`
		Catalogname3       interface{} `json:"catalogname3"`
		Commonname         interface{} `json:"commonname"`
		Matrial            interface{} `json:"matrial"`
		Characteristic     interface{} `json:"characteristic"`
		RelationId         string      `json:"relationId"`
		CatalogCode        string      `json:"catalogCode"`
		SpecificationCode  string      `json:"specificationCode"`
		ProductionCode     string      `json:"productionCode"`
		Regcardid          string      `json:"regcardid"`
		Regcardnm          string      `json:"regcardnm"`
		RegcardName        string      `json:"regcardName"`
		Productid          string      `json:"productid"`
		ProductName        string      `json:"productName"`
		Goodsid            string      `json:"goodsid"`
		Specification      string      `json:"specification"`
		Model              string      `json:"model"`
		IsUsing            int         `json:"isUsing"`
		RelationStatus     string      `json:"relationStatus"`
		AddUserId          string      `json:"addUserId"`
		AddUserName        string      `json:"addUserName"`
		AddTime            time.Time   `json:"addTime"`
		LastUpdateUserId   string      `json:"lastUpdateUserId"`
		LastUpdateUserName string      `json:"lastUpdateUserName"`
		LastUpdateTime     time.Time   `json:"lastUpdateTime"`
		AuditRemark        interface{} `json:"auditRemark"`
		AuditUserId        interface{} `json:"auditUserId"`
		AuditUserName      interface{} `json:"auditUserName"`
		AuditTime          interface{} `json:"auditTime"`
		UdiCode            interface{} `json:"udiCode"`
		GgxhCode           string      `json:"ggxhCode"`
		Oldregcardnm       *string     `json:"oldregcardnm"`
		MapingCode         interface{} `json:"mapingCode"`
		CodeOld            interface{} `json:"codeOld"`
		CodeShow           string      `json:"codeShow"`
		DataType           interface{} `json:"dataType"`
		Registrant         string      `json:"registrant"`
	} `json:"rows"`
	Page        int         `json:"page"`
	Count       int         `json:"count"`
	FirstResult int         `json:"firstResult"`
	MaxResults  int         `json:"maxResults"`
	Success     bool        `json:"success"`
	Result      interface{} `json:"result"`
	Conditions  struct {
		ProductionCode    string `json:"productionCode"`
		SpecificationCode string `json:"specificationCode"`
	} `json:"conditions"`
	Msg       string      `json:"msg"`
	Form      interface{} `json:"form"`
	Code      int         `json:"code"`
	OperCount int         `json:"operCount"`
	Sord      string      `json:"sord"`
	Sidx      string      `json:"sidx"`
	Orderby   interface{} `json:"orderby"`
}

func GetDate1004(page int) {

	url := "https://code.nhsa.gov.cn/hc/stdSpecification/getStdSpecificationListDataCompanyReport.html"
	method := "POST"

	text := "_search=湘械注准20202140143&nd=&rows=50&page=" + strconv.Itoa(page) + "&sidx=&sord=desc"

	payload := strings.NewReader(text)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", "__jsluid_s=9d335ba068cdccfdba87d26868685ddd; queryCondition=9be8ff2ed273bc199e6707a6822f559f%3D%7B%22specificationCode%22%3A%22%22%2C%22commonname%22%3A%22%22%2C%22companyName%22%3A%22%22%2C%22catalogname1%22%3A%22%22%2C%22catalogname2%22%3A%22%22%2C%22catalogname3%22%3A%22%22%2C%22regcardNm%22%3A%22%22%2C%22releaseVersion%22%3A%22%22%7D; JSESSIONID=7382B3D591669341EC48AB4A98CDEAF7; JSESSIONID=C620AEE883E894EFD6756E33A4BF49A2")
	req.Header.Add("Origin", "https://code.nhsa.gov.cn")
	req.Header.Add("Referer", "https://code.nhsa.gov.cn/hc/stdSpecification/toStdSpecificationCompanyReportList.html")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var doc1104 Doc1104
	json.Unmarshal(body, &doc1104)
	for i, row := range doc1104.Rows {
		getDataInfo1104(row.SpecificationCode, page, i)
	}

}

func getDataInfo1104(code string, page, num int) {

	url := "https://code.nhsa.gov.cn/hc/stdYgbData/getPublicHcDataList.html"
	method := "POST"

	text := "productid=&regcardid=&releaseVersion=&productNameHide=&regcardnm=&productName=&specification=&model=&specificationCode=" + code +
		"&_search=false&nd=&rows=100000&page=1&sidx=&sord=desc"

	payload := strings.NewReader(text)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", "__jsluid_s=9d335ba068cdccfdba87d26868685ddd; JSESSIONID=DE21880097ECE4D818797C119191AAE5; queryCondition=6560134ae133bc837036b4198b79e7eb%3D%7B%22productid%22%3A%22%22%2C%22regcardid%22%3A%22%22%2C%22releaseVersion%22%3A%22%22%2C%22productNameHide%22%3A%22%22%2C%22regcardnm%22%3A%22%22%2C%22productName%22%3A%22%22%2C%22specification%22%3A%22%22%2C%22model%22%3A%22%22%7D; JSESSIONID=C620AEE883E894EFD6756E33A4BF49A2")
	req.Header.Add("Origin", "https://code.nhsa.gov.cn")
	req.Header.Add("Referer", "https://code.nhsa.gov.cn/hc/stdYgbData/toHcList2.html?specificationCode=C1402022620000904065&releaseVersion=20220707&catalogname=%E5%9F%BA%E7%A1%80%E5%8D%AB%E7%94%9F%E6%9D%90%E6%96%99/%E5%B8%B8%E8%A7%84%E5%8C%BB%E7%96%97%E7%94%A8%E5%93%81/%E5%85%B6%E4%BB%96%E5%B8%B8%E8%A7%84%E5%8C%BB%E7%96%97%E7%94%A8%E5%93%81/%E5%8F%A3%E7%BD%A9/%E4%B8%8D%E5%8C%BA%E5%88%86/%E5%8F%A3%E7%BD%A9&companyName=%E5%8F%AF%E5%AD%9A%E5%8C%BB%E7%96%97%E7%A7%91%E6%8A%80%E8%82%A1%E4%BB%BD%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

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

	db, _ := db.MysqlConfigure()

	var dao1104 Dao1104
	json.Unmarshal(body, &dao1104)
	for i, row := range dao1104.Rows {
		fmt.Printf("\n页码:%d\tnum:%d\t个数:%d\tcode:%s", page, num, i, code)
		db.Table("t_medicine_1104").Create(&row)
	}
}
