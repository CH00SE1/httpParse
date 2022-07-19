package guojiayibao

import (
	"gorm.io/gorm"
	"time"
)

/**
 * @title 对象
 * @author xiongshao
 * @date 2022-07-19 08:07:16
 */

// 1101
type Medicine1001 struct {
	gorm.Model
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
}

// 1102
type Medicine1002 struct {
	gorm.Model
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
}

// 1103
type Medicine1003 struct {
	gorm.Model
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
}
