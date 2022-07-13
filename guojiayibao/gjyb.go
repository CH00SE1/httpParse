package guojiayibao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
 * @title 国家医保
 * @author xiongshao
 * @date 2022-06-24 12:26:47
 */

// 返回对象1003
type GjybInfo struct {
	Records int `json:"records"`
	Total   int `json:"total"`
	Rows    []struct {
		MaterialName            string `json:"materialName"`
		CompanyNameSc           string `json:"companyNameSc"`
		RegisteredProductName   string `json:"registeredProductName"`
		Unit                    string `json:"unit"`
		ApprovalCode            string `json:"approvalCode"`
		RegisteredOutlook       string `json:"registeredOutlook"`
		RegisteredMedicinemodel string `json:"registeredMedicinemodel"`
		GoodsStandardCode       string `json:"goodsStandardCode"`
		GoodsCode               string `json:"goodsCode"`
		MinUnit                 string `json:"minUnit"`
		Factor                  int    `json:"factor"`
		GoodsName               string `json:"goodsName"`
		ProductRemark           string `json:"productRemark,omitempty"`
		ProductName             string `json:"productName,omitempty"`
	} `json:"rows"`
	Page        int         `json:"page"`
	Count       int         `json:"count"`
	FirstResult int         `json:"firstResult"`
	MaxResults  int         `json:"maxResults"`
	Success     bool        `json:"success"`
	Result      interface{} `json:"result"`
	Conditions  struct {
	} `json:"conditions"`
	Msg       interface{} `json:"msg"`
	Form      interface{} `json:"form"`
	Code      int         `json:"code"`
	OperCount int         `json:"operCount"`
	Sord      string      `json:"sord"`
	Sidx      string      `json:"sidx"`
	Orderby   interface{} `json:"orderby"`
}

// 1001返回对象
type Drug1001 struct {
	Records int `json:"records"`
	Total   int `json:"total"`
	Rows    []struct {
		PiecesId                     string    `json:"piecesId"`
		PiecesCode                   string    `json:"piecesCode"`
		PiecesName                   string    `json:"piecesName"`
		MedicinalMaterialsName       string    `json:"medicinalMaterialsName"`
		ProcessingMethod             string    `json:"processingMethod"`
		EfficacyClassification       string    `json:"efficacyClassification"`
		SubjectSource                string    `json:"subjectSource"`
		MaterialSource               string    `json:"materialSource"`
		PharmaceuticalSite           string    `json:"pharmaceuticalSite"`
		PropertyflavorChanneltropism string    `json:"propertyflavorChanneltropism"`
		FunctionIndications          string    `json:"functionIndications"`
		UsageDosage                  string    `json:"usageDosage"`
		PaymentPolicy                string    `json:"paymentPolicy"`
		PaymentPolicyP               string    `json:"paymentPolicyP"`
		SpecificationName            string    `json:"specificationName"`
		SpecificationPageNumber      string    `json:"specificationPageNumber"`
		SpecificationAttachment      string    `json:"specificationAttachment"`
		AreaName                     string    `json:"areaName"`
		AreaId                       string    `json:"areaId"`
		InitializationState          int       `json:"initializationState"`
		SubmitTime                   time.Time `json:"submitTime"`
		IsUsing                      int       `json:"isUsing"`
		IsUsingRemark                string    `json:"isUsingRemark"`
		ReauditUserId                string    `json:"reauditUserId"`
		ReauditUserName              string    `json:"reauditUserName"`
		ReauditAddTime               time.Time `json:"reauditAddTime"`
		ReauditRemark                *string   `json:"reauditRemark"`
		AuditUserId                  string    `json:"auditUserId"`
		AuditUserName                string    `json:"auditUserName"`
		AuditAddTime                 time.Time `json:"auditAddTime"`
		AuditRemark                  *string   `json:"auditRemark"`
		AddUserId                    string    `json:"addUserId"`
		AddUserName                  string    `json:"addUserName"`
		AddTime                      time.Time `json:"addTime"`
		LastUpdateUserId             string    `json:"lastUpdateUserId"`
		LastUpdateUserName           string    `json:"lastUpdateUserName"`
		LastUpdateTime               time.Time `json:"lastUpdateTime"`
		EfficacyClassificationid     string    `json:"efficacyClassificationid"`
		IsReport                     string    `json:"isReport"`
		RkFlag                       bool      `json:"rkFlag"`
		RkkFlag                      bool      `json:"rkkFlag"`
		SourceAreaName               string    `json:"sourceAreaName"`
		SourceAreaId                 string    `json:"sourceAreaId"`
		PushVersion                  string    `json:"pushVersion"`
		FinalPushStatus              int       `json:"finalPushStatus"`
		Message                      string    `json:"message"`
		IsuFlag                      int       `json:"isuFlag"`
		RkTime                       time.Time `json:"rkTime"`
		PushTime                     time.Time `json:"pushTime"`
	} `json:"rows"`
	Page        int         `json:"page"`
	Count       int         `json:"count"`
	FirstResult int         `json:"firstResult"`
	MaxResults  int         `json:"maxResults"`
	Success     bool        `json:"success"`
	Result      interface{} `json:"result"`
	Conditions  struct {
	} `json:"conditions"`
	Msg       interface{} `json:"msg"`
	Form      interface{} `json:"form"`
	Code      int         `json:"code"`
	OperCount int         `json:"operCount"`
	Sord      string      `json:"sord"`
	Sidx      string      `json:"sidx"`
	Orderby   interface{} `json:"orderby"`
}

type DurgInfo1002 struct {
	Records int `json:"records"`
	Total   int `json:"total"`
	Rows    []struct {
		PreparationId                    string     `json:"preparationId"`
		PreparationCode                  string     `json:"preparationCode"`
		PreparationType                  string     `json:"preparationType"`
		PreparationPrename               string     `json:"preparationPrename"`
		PreparationMedicinemodel         string     `json:"preparationMedicinemodel"`
		PreparationOutlook               string     `json:"preparationOutlook"`
		PreparationFactor                string     `json:"preparationFactor"`
		PreparationPacknuit              string     `json:"preparationPacknuit"`
		PreparationNuit                  string     `json:"preparationNuit"`
		PreparationMaterialname          string     `json:"preparationMaterialname"`
		PreparationName                  string     `json:"preparationName"`
		PreparationAddress               string     `json:"preparationAddress"`
		PreparationCommissionname        string     `json:"preparationCommissionname"`
		PreparationCommissionaddress     string     `json:"preparationCommissionaddress"`
		PreparationApprovalcode          string     `json:"preparationApprovalcode"`
		PreparationValiditydate          *time.Time `json:"preparationValiditydate"`
		PreparationPermitnumber          string     `json:"preparationPermitnumber"`
		PreparationExestandard           string     `json:"preparationExestandard"`
		PreparationApplicabledisease     string     `json:"preparationApplicabledisease"`
		PreparationDosage                string     `json:"preparationDosage"`
		PreparationChildmedication       string     `json:"preparationChildmedication"`
		PreparationOldatientmedication   string     `json:"preparationOldatientmedication"`
		PreparationContactname           string     `json:"preparationContactname"`
		PreparationContactnumber         string     `json:"preparationContactnumber"`
		PreparationPerdocattachment      string     `json:"preparationPerdocattachment"`
		PreparationApprovaldocattachment string     `json:"preparationApprovaldocattachment"`
		PreparationDesdocattachment      string     `json:"preparationDesdocattachment"`
		AreaName                         string     `json:"areaName"`
		AreaId                           string     `json:"areaId"`
		HosId                            string     `json:"hosId"`
		HosName                          string     `json:"hosName"`
		InitializationState              int        `json:"initializationState"`
		SubmitTime                       time.Time  `json:"submitTime"`
		ReauditUserId                    string     `json:"reauditUserId"`
		ReauditUserName                  string     `json:"reauditUserName"`
		ReauditAddTime                   time.Time  `json:"reauditAddTime"`
		ReauditRemark                    string     `json:"reauditRemark"`
		AuditUserId                      string     `json:"auditUserId"`
		AuditUserName                    string     `json:"auditUserName"`
		AuditAddTime                     time.Time  `json:"auditAddTime"`
		AuditRemark                      string     `json:"auditRemark"`
		AddUserId                        string     `json:"addUserId"`
		AddUserName                      string     `json:"addUserName"`
		AddTime                          time.Time  `json:"addTime"`
		LastUpdateUserId                 string     `json:"lastUpdateUserId"`
		LastUpdateUserName               string     `json:"lastUpdateUserName"`
		LastUpdateTime                   time.Time  `json:"lastUpdateTime"`
		RkFlag                           bool       `json:"rkFlag"`
		RkkFlag                          bool       `json:"rkkFlag"`
		FinalPushStatus                  int        `json:"finalPushStatus"`
		PushVersion                      string     `json:"pushVersion"`
		Message                          string     `json:"message"`
		DataType                         int        `json:"dataType"`
		IsuFlag                          int        `json:"isuFlag"`
		RkTime                           time.Time  `json:"rkTime"`
		PushTime                         time.Time  `json:"pushTime"`
		ProductInsuranceType             int        `json:"productInsuranceType"`
		ProductName                      string     `json:"productName"`
		ProductMedicineModel             string     `json:"productMedicineModel"`
		ProductRemark                    string     `json:"productRemark"`
		ProductCode                      string     `json:"productCode"`
		PayStandard                      string     `json:"payStandard"`
	} `json:"rows"`
	Page        int         `json:"page"`
	Count       int         `json:"count"`
	FirstResult int         `json:"firstResult"`
	MaxResults  int         `json:"maxResults"`
	Success     bool        `json:"success"`
	Result      interface{} `json:"result"`
	Conditions  struct {
	} `json:"conditions"`
	Msg       interface{} `json:"msg"`
	Form      interface{} `json:"form"`
	Code      int         `json:"code"`
	OperCount int         `json:"operCount"`
	Sord      string      `json:"sord"`
	Sidx      string      `json:"sidx"`
	Orderby   interface{} `json:"orderby"`
}

func GetData1003(page int) GjybInfo {

	url := "https://code.nhsa.gov.cn/yp/stdGoodsPublic/getStdGoodsPublicData.html"
	method := "POST"

	text := "goodsCode=&companyNameSc=&registeredProductName=&approvalCode=&_search=false&nd=1657678747531&rows=10000&page=" + strconv.Itoa(page) + "&sidx=&sord=asc"

	payload := strings.NewReader(text)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", "__jsluid_s=432f80e073d0ca463ea5c40138b6347d; JSESSIONID=302417BF25577AEA2EA7252027C7C02B; queryCondition=c56c3a956c2b31c025f1905cf3b1cc3a%3D%7B%22goodsCode%22%3A%22%22%2C%22companyNameSc%22%3A%22%22%2C%22registeredProductName%22%3A%22%22%2C%22approvalCode%22%3A%22%E5%9B%BD%E8%8D%AF%E5%87%86%E5%AD%97H20180004%22%7D; JSESSIONID=A5738BC156E3C9100E2DE42B3BFC1B62")
	req.Header.Add("Origin", "https://code.nhsa.gov.cn")
	req.Header.Add("Referer", "https://code.nhsa.gov.cn/yp/toPublicList3.html")
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
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected EOF") && len(body) != 0 {
			log.Fatal(err, page)
		}
	}
	var gjybInfo GjybInfo
	json.Unmarshal(body, &gjybInfo)
	return gjybInfo
}

func GetData1001(page int) Drug1001 {

	url := "https://code.nhsa.gov.cn/yp/stdChineseMedicinalDecoctionPieces/getPiecesRkData.html"
	method := "POST"

	text := "piecesCode=&piecesName=&_search=false&nd=1657674738005&rows=100&page=" + strconv.Itoa(page) + "&sidx=&sord=asc"

	payload := strings.NewReader(text)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println("request error:", err)
	}
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", "__jsluid_s=9d335ba068cdccfdba87d26868685ddd; pageSelect=40ca710142ff03a3617cc9ff170a5651%3D1; JSESSIONID=1ACDC582854E0787C527E24F496FFF2C; queryCondition=575f478bdd1b6665386a42ec3cf354b2%3D%7B%22piecesCode%22%3A%22%22%2C%22piecesName%22%3A%22%E5%BD%93%E5%BD%92%22%7D; JSESSIONID=BEB1686381D3EC1323765305B5E1785A")
	req.Header.Add("Origin", "https://code.nhsa.gov.cn")
	req.Header.Add("Referer", "https://code.nhsa.gov.cn/yp/toRkList.html")
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
		fmt.Println("response error:", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ioutil error:", err)
	}
	var drug1001 Drug1001
	json.Unmarshal(body, &drug1001)
	return drug1001
}

func GetData1002(page int) DurgInfo1002 {

	url := "https://code.nhsa.gov.cn/yp/stdChineseMedicinalDecoctionPieces/getYnzjHospreparationRkData.html"
	method := "POST"

	text := "hosName=&preparationCode=&preparationPrename=&preparationApprovalcode=&_search=false&nd=1656137002632&rows=1000&page=" + strconv.Itoa(page) + "&sidx=&sord=asc"

	payload := strings.NewReader(text)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return DurgInfo1002{}
	}
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", "__jsluid_s=432f80e073d0ca463ea5c40138b6347d; pageSelect=0651c039aeb194b9cc7f459fd752d7f1%3D1; queryCondition=17fdc4365e50daabc7f92a2a05ee0ed0%3D%7B%22hosName%22%3A%22%E6%B9%98%E9%9B%85%22%2C%22preparationCode%22%3A%22%22%2C%22preparationPrename%22%3A%22%22%2C%22preparationApprovalcode%22%3A%22%22%7D; JSESSIONID=29F03F078AD61679A8EBE5E98761CE87; JSESSIONID=2B9DF6843058E832395D7084B337B1D1")
	req.Header.Add("Origin", "https://code.nhsa.gov.cn")
	req.Header.Add("Referer", "https://code.nhsa.gov.cn/yp/toRkList2.html")
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
		return DurgInfo1002{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return DurgInfo1002{}
	}
	var durgInfo1002 DurgInfo1002
	json.Unmarshal(body, &durgInfo1002)
	fmt.Println(page, "------", text)
	return durgInfo1002
}
