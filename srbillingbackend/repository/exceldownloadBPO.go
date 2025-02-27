package repository

import (
	"database/sql"
	"fmt"

	"github.com/Naveeshkumar24/internal/models"
	"github.com/xuri/excelize/v2"
)

type ExcelDownloadBPO struct {
	db *sql.DB
}

func NewExcelDownloadCPO(db *sql.DB) *ExcelDownloadBPO {
	return &ExcelDownloadBPO{db: db}
}

func (e *ExcelDownloadBPO) FetchExcelCPO() ([]models.ExcelDownloadBPO, error) {
	var data []models.ExcelDownloadBPO

	if e.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	rows, err := e.db.Query("SELECT * FROM billingposubmitteddata")
	if err != nil {
		fmt.Println("Database query error:", err)
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var excelData models.ExcelDownloadBPO
		if err := rows.Scan(
			&excelData.ID,
			&excelData.Timestamp,
			&excelData.EnggName,
			&excelData.Supplier,
			&excelData.BillNo,
			&excelData.BillDate,
			&excelData.CustomerName,
			&excelData.CustomerPoNo,
			&excelData.CustomerPoDate,
			&excelData.ItemDescription,
			&excelData.BilledQty,
			&excelData.Unit,
			&excelData.NetValue,
			&excelData.CGST,
			&excelData.IGST,
			&excelData.Totaltax,
			&excelData.Gross,
			&excelData.DispatchThrough,
		); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		data = append(data, excelData)
	}
	if len(data) == 0 {
		fmt.Println("No data found in customerposubmitteddata table")
		return []models.ExcelDownloadBPO{}, nil
	}
	fmt.Printf("Fetched %d records\n", len(data))
	return data, nil
}

func (e *ExcelDownloadBPO) CreateExcelDownloadCPO() (*excelize.File, error) {
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	data, err := e.FetchExcelCPO()
	if err != nil {
		return nil, err
	}
	sheetName := "CustomerPO"
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
	return file, nil
}
