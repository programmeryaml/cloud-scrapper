package alicloud

// InstanceInfo represents the simplified structure for ECS instance information
type PriceInfo struct {
	InstanceID        string  `json:"InstanceID" xml:"InstanceID"`
	InvoiceDiscount   float64 `json:"InvoiceDiscount" xml:"InvoiceDiscount"`
	PretaxGrossAmount float64 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount"`
	Item              string  `json:"Item" xml:"Item"`
	InstanceId        string  `json:"InstanceId" xml:"InstanceId"`
	OutstandingAmount float64 `json:"OutstandingAmount" xml:"OutstandingAmount"`
	ProductCode       string  `json:"ProductCode" xml:"ProductCode"`
	Region            string  `json:"Region" xml:"Region"`
	InstanceSpec      string  `json:"InstanceSpec" xml:"InstanceSpec"`
	BillingType       string  `json:"BillingType" xml:"BillingType"`
	Tag               string  `json:"Tag" xml:"Tag"`
	NickName          string  `json:"NickName" xml:"NickName"`
	Currency          string  `json:"Currency" xml:"Currency"`
	DeductedByCoupons float64 `json:"DeductedByCoupons" xml:"DeductedByCoupons"`
	Zone              string  `json:"Zone" xml:"Zone"`
	PretaxAmount      float64 `json:"PretaxAmount" xml:"PretaxAmount"`
	IntranetIP        string  `json:"IntranetIP" xml:"IntranetIP"`
}
