package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/Naveeshkumar24/internal/models"
)

type Query struct {
	db   *sql.DB
	Time *time.Location
}

func NewQuery(db *sql.DB) *Query {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatalf("Failed to load time zone: %v", err)
	}

	return &Query{
		db:   db,
		Time: loc,
	}
}
func (q *Query) CreateTables() error {
	tx, err := q.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queries := []string{
		`CREATE TABLE IF NOT EXISTS customername (
			customer_name VARCHAR(255) NOT NULL UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS unit (	
			unit_name VARCHAR(100) NOT NULL UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS enggname(
			engg_name VARCHAR(100) NOT NULL UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS supplier(
			supplier_name VARCHAR(100) NOT NULL UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS billingposubmitteddata (
			id SERIAL PRIMARY KEY,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			engg_name VARCHAR(255),
			supplier VARCHAR(255),
			bill_no VARCHAR(100) NOT NULL UNIQUE,
			bill_date DATE,
			customer_name VARCHAR(255) ,
			customer_po_no VARCHAR(100),
			customer_po_date DATE,
			item_description TEXT,
			billed_qty INT,
			unit VARCHAR(100) ,
			net_value DECIMAL(12,2),
			cgst DECIMAL(12,2),
			igst DECIMAL(12,2),
			total_tax DECIMAL(12,2),
			gross DECIMAL(12,2),
			dispatch_through VARCHAR(255)
		)`,
	}

	for _, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			log.Printf("Failed to execute query: %s", query)
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	log.Println("All tables created successfully.")
	return nil
}

func (q *Query) FetchDropDown() ([]models.BillingPoDropDown, error) {
	var dropdowns []models.BillingPoDropDown

	rows, err := q.db.Query(`
		SELECT e.engg_name,s.supplier_name,c.customer_name,u.unit_name
		FROM enggname e
		CROSS JOIN supplier s
		CROSS JOIN customername c
		CROSS JOIN unit u;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dropdown models.BillingPoDropDown
		if err := rows.Scan(&dropdown.EnggName, &dropdown.Supplier, &dropdown.CustomerName, &dropdown.Unit); err != nil {
			return nil, err
		}
		dropdowns = append(dropdowns, dropdown)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dropdowns, nil
}

func (q *Query) SubmitFormBillingPoData(data models.BillingPo) error {
	_, err := q.db.Exec(`
		INSERT INTO  billingposubmitteddata (
			engg_name, supplier, bill_no, bill_date, customer_name,
			customer_po_no, customer_po_date, item_description, billed_qty, unit, net_value,
			cgst, igst, total_tax, gross, dispatch_through
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`,
		data.EnggName, data.Supplier, data.BillNo, data.BillDate,
		data.CustomerName, data.CustomerPoNo, data.CustomerPoDate, data.ItemDescription,
		data.BilledQty, data.Unit, data.NetValue, data.CGST, data.IGST, data.Totaltax,
		data.Gross, data.DispatchThrough,
	)

	if err != nil {
		log.Printf("Failed to insert BillingPo data: %v", err)
		return err
	}

	log.Println("BillingPo data submitted successfully.")
	return nil
}

func (q *Query) FetchBillingPoData() ([]models.BillingPo, error) {
	var billingPoList []models.BillingPo

	rows, err := q.db.Query(`
		SELECT id, timestamp, engg_name, supplier, bill_no, bill_date, customer_name, 
		       customer_po_no, customer_po_date, item_description, billed_qty, unit, net_value, 
		       cgst, igst, total_tax, gross, dispatch_through 
		FROM  billingposubmitteddata;
	`)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var billingPo models.BillingPo
		err := rows.Scan(
			&billingPo.ID, &billingPo.Timestamp, &billingPo.EnggName, &billingPo.Supplier,
			&billingPo.BillNo, &billingPo.BillDate, &billingPo.CustomerName, &billingPo.CustomerPoNo,
			&billingPo.CustomerPoDate, &billingPo.ItemDescription, &billingPo.BilledQty,
			&billingPo.Unit, &billingPo.NetValue, &billingPo.CGST, &billingPo.IGST,
			&billingPo.Totaltax, &billingPo.Gross, &billingPo.DispatchThrough,
		)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, err
		}
		billingPoList = append(billingPoList, billingPo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("Fetched %d BillingPo records.", len(billingPoList))
	return billingPoList, nil
}

func (q *Query) UpdateBillingPoData(data models.BillingPo) error {
	_, err := q.db.Exec(`
		UPDATE billingposubmitteddata SET
			engg_name= $1, supplier = $2, bill_no = $3, bill_date = $4, 
			customer_name = $5, customer_po_no = $6, customer_po_date = $7, item_description = $8, 
			billed_qty = $9, unit = $10, net_value = $11, cgst = $12, igst = $13, total_tax = $14, 
			gross = $15, dispatch_through = $16
		WHERE id = $17`,
		data.EnggName, data.Supplier, data.BillNo, data.BillDate,
		data.CustomerName, data.CustomerPoNo, data.CustomerPoDate, data.ItemDescription,
		data.BilledQty, data.Unit, data.NetValue, data.CGST, data.IGST, data.Totaltax,
		data.Gross, data.DispatchThrough, data.ID,
	)

	if err != nil {
		log.Printf("Failed to update BillingPo data for ID %d: %v", data.ID, err)
		return err
	}

	log.Printf("BillingPo data updated successfully for ID %d.", data.ID)
	return nil
}
func (q *Query) DeleteBillingPoData(id int) error {
	_, err := q.db.Exec(`DELETE FROM billingposubmitteddata WHERE id = $1`, id)
	if err != nil {
		log.Printf("Failed to delete BillingPo data for ID %d: %v", id, err)
		return err
	}

	log.Printf("BillingPo data deleted successfully for ID %d.", id)
	return nil
}
