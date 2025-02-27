package models

import (
	"io"

	"github.com/xuri/excelize/v2"
)

type ExcelDownloadBPO struct {
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
type ExcelDownloadBPOInterface interface {
	CreateExcelDownloadCPO(*io.ReadCloser) (*excelize.File, error)
}
