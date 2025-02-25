package models

import "net/http"

type BillingPo struct {
	ID              int     `json:"id"`
	Timestamp       string  `json:"timestamp"`
	EnggName        string  `json:"Engg_Name"`
	Supplier        string  `json:"Supplier"`
	BillNo          string  `json:"Bill_No"`
	BillDate        string  `json:"Bill_Date"`
	CustomerName    string  `json:"Customer_Name"`
	CustomerPoNo    string  `json:"Customer_Po_No"`
	CustomerPoDate  string  `json:"Customer_Po_Date"`
	ItemDescription string  `json:"Item_Description"`
	BilledQty       int     `json:"Billed_Qty"`
	Unit            string  `json:"unit"`
	NetValue        float64 `json:"Net_Value"`
	CGST            float64 `json:"CGST"`
	IGST            float64 `json:"IGST"`
	Totaltax        float64 `json:"Total_tax"`
	Gross           float64 `json:"Gross"`
	DispatchThrough string  `json:"Dispatch_Through"`
}
type BillingPoDropDown struct {
	EnggName     string `json:"Engg_Name"`
	Supplier     string `json:"Supplier"`
	CustomerName string `json:"Customer_Name"`
	Unit         string `json:"unit"`
}

type CustomerPoInterface interface {
	FetchDropDown() ([]BillingPoDropDown, error)
	SubmitFormCustomerPoData(billingPo BillingPo) error
	FetchCustomerPoData(r *http.Request) ([]BillingPo, error)
	UpdateCustomerPoData(customerPo BillingPo) error
}
