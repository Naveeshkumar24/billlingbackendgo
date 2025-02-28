package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/naveeshkumar24/internal/models"
	"github.com/xuri/excelize/v2"
)

type ExcelDownloadBPOHandler struct {
	bpoRepo models.BillingPoInterface
}

func NewExcelDownloadBPOHandler(bpoRepo models.BillingPoInterface) *ExcelDownloadBPOHandler {
	return &ExcelDownloadBPOHandler{
		bpoRepo: bpoRepo,
	}
}

func (edb *ExcelDownloadBPOHandler) DownloadBPO(w http.ResponseWriter, r *http.Request) {
	data, err := edb.bpoRepo.FetchBillingPoData(r)
	if err != nil {
		http.Error(w, "Failed to fetch BPO data", http.StatusInternalServerError)
		return
	}
	file := excelize.NewFile()
	sheetName := "BPO"
	file.NewSheet(sheetName)

	headers := []string{
		"ID", "Timestamp", "Engineer Name", "Supplier", "Bill No", "Bill Date", "Customer Name",
		"Customer PO No", "Customer PO Date", "Item Description", "Billed Qty", "Unit",
		"Net Value", "CGST", "IGST", "Total Tax", "Gross", "Dispatch Through",
	}

	for colIndex, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(colIndex+1, 1)
		file.SetCellValue(sheetName, cell, header)
	}

	file.DeleteSheet("Sheet1")

	for i, record := range data {
		rowNum := i + 2
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", rowNum), record.ID)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", rowNum), record.Timestamp)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", rowNum), record.EnggName)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", rowNum), record.Supplier)
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", rowNum), record.BillNo)
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", rowNum), record.BillDate)
		file.SetCellValue(sheetName, fmt.Sprintf("G%d", rowNum), record.CustomerName)
		file.SetCellValue(sheetName, fmt.Sprintf("H%d", rowNum), record.CustomerPoNo)
		file.SetCellValue(sheetName, fmt.Sprintf("I%d", rowNum), record.CustomerPoDate)
		file.SetCellValue(sheetName, fmt.Sprintf("J%d", rowNum), record.ItemDescription)
		file.SetCellValue(sheetName, fmt.Sprintf("K%d", rowNum), record.BilledQty)
		file.SetCellValue(sheetName, fmt.Sprintf("L%d", rowNum), record.Unit)
		file.SetCellValue(sheetName, fmt.Sprintf("M%d", rowNum), record.NetValue)
		file.SetCellValue(sheetName, fmt.Sprintf("N%d", rowNum), record.CGST)
		file.SetCellValue(sheetName, fmt.Sprintf("O%d", rowNum), record.IGST)
		file.SetCellValue(sheetName, fmt.Sprintf("P%d", rowNum), record.Totaltax)
		file.SetCellValue(sheetName, fmt.Sprintf("Q%d", rowNum), record.Gross)
		file.SetCellValue(sheetName, fmt.Sprintf("R%d", rowNum), record.DispatchThrough)
	}

	tempDir := "/tmp"
	if os.Getenv("OS") == "Windows_NT" {
		tempDir = os.Getenv("TEMP")
	}
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create temporary directory", http.StatusInternalServerError)
		return
	}

	filepath := fmt.Sprintf("%s/bpodata.xlsx", tempDir)
	if err := file.SaveAs(filepath); err != nil {
		http.Error(w, "Failed to save Excel file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=bpodata.xlsx")
	http.ServeFile(w, r, filepath)
}
