package repository

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Naveeshkumar24/internal/models"
	"github.com/Naveeshkumar24/pkg/database"
)

type BillingPoRepository struct {
	db *sql.DB
}

func NewBillingPoRepository(db *sql.DB) *BillingPoRepository {
	return &BillingPoRepository{
		db: db,
	}
}
func (c *BillingPoRepository) FetchDropDown() ([]models.BillingPoDropDown, error) {
	query := database.NewQuery(c.db)
	res, err := query.FetchDropDown()
	if err != nil {
		log.Printf("Database query failed: %v", err)
		return nil, err
	}
	if len(res) == 0 {
		log.Println("No data found in FetchDropDown query")
		return nil, sql.ErrNoRows
	}
	log.Println("Successfully fetched dropdown data")
	return res, nil
}
func (c *BillingPoRepository) SubmitFormBillingPoData(data models.BillingPo) error {
	query := database.NewQuery(c.db)
	err := query.SubmitFormBillingPoData(data)
	if err != nil {
		log.Printf("Failed to submit customer PO data: %v", err)
		return err
	}
	return nil
}
func (c *BillingPoRepository) FetchBillingPoData(r *http.Request) ([]models.BillingPo, error) {
	query := database.NewQuery(c.db)
	res, err := query.FetchBillingPoData()
	if err != nil {
		log.Printf("Failed to fetch customer PO data: %v", err)
		return nil, err
	}
	return res, nil
}
func (c *BillingPoRepository) UpdateBillingPoData(data models.BillingPo) error {
	query := database.NewQuery(c.db)
	err := query.UpdateBillingPoData(data)
	if err != nil {
		log.Printf("Failed to update customer PO data: %v", err)
		return err
	}

	return nil
}

func (c *BillingPoRepository) DeleteBillingPo(id int) error {
	tx, err := c.db.Begin()
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec("DELETE FROM customerposubmitteddata WHERE id = $1", id)
	if err != nil {
		log.Printf("Failed to delete record with id %d: %v", id, err)
		return err
	}

	return nil
}
