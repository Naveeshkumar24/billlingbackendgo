package models

import "net/http"

type BillingPo struct {
	ID              int     `json:"id"`
	Timestamp       string  `json:"timestamp"`
	EnggName        string  `json:"engg_Name"`
	Supplier        string  `json:"supplier"`
	BillNo          string  `json:"bill_No"`
	BillDate        string  `json:"bill_Date"`
	CustomerName    string  `json:"customer_Name"`
	CustomerPoNo    string  `json:"customer_Po_No"`
	CustomerPoDate  string  `json:"customer_Po_Date"`
	ItemDescription string  `json:"item_Description"`
	BilledQty       int     `json:"billed_Qty"`
	Unit            string  `json:"unit"`
	NetValue        float64 `json:"net_Value"`
	CGST            float64 `json:"cgst"`
	IGST            float64 `json:"igst"`
	Totaltax        float64 `json:"total_tax"`
	Gross           float64 `json:"gross"`
	DispatchThrough string  `json:"dispatch_through"`
}
type BillingPoDropDown struct {
	EnggName     string `json:"engg_name"`
	Supplier     string `json:"supplier_name"`
	CustomerName string `json:"customer_name"`
	Unit         string `json:"unit_name"`
}

type BillingPoInterface interface {
	FetchDropDown() ([]BillingPoDropDown, error)
	SubmitFormCustomerPoData(billingPo BillingPo) error
	FetchCustomerPoData(r *http.Request) ([]BillingPo, error)
	UpdateCustomerPoData(customerPo BillingPo) error
}
