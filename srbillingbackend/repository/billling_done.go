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

// FetchCustomerPoData implements models.BillingPoInterface.
func (b *BillingPoRepository) FetchCustomerPoData(r *http.Request) ([]models.BillingPo, error) {
	panic("unimplemented")
}

// SubmitFormCustomerPoData implements models.BillingPoInterface.
func (b *BillingPoRepository) SubmitFormCustomerPoData(billingPo models.BillingPo) error {
	panic("unimplemented")
}

// UpdateCustomerPoData implements models.BillingPoInterface.
func (b *BillingPoRepository) UpdateCustomerPoData(customerPo models.BillingPo) error {
	panic("unimplemented")
}

func NewBillingPoRepository(db *sql.DB) *BillingPoRepository {
	return &BillingPoRepository{
		db: db,
	}
}

func (b *BillingPoRepository) FetchDropDown() ([]models.BillingPoDropDown, error) {
	query := database.NewQuery(b.db)
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

func (b *BillingPoRepository) SubmitFormBillingPoData(data models.BillingPo) error {
	query := database.NewQuery(b.db)
	err := query.SubmitFormBillingPoData(data)
	if err != nil {
		log.Printf("Failed to submit billing PO data: %v", err)
		return err
	}
	return nil
}

func (b *BillingPoRepository) FetchBillingPoData(r *http.Request) ([]models.BillingPo, error) {
	query := database.NewQuery(b.db)
	res, err := query.FetchBillingPoData()
	if err != nil {
		log.Printf("Failed to fetch billing PO data: %v", err)
		return nil, err
	}
	return res, nil
}

func (b *BillingPoRepository) UpdateBillingPoData(data models.BillingPo) error {
	query := database.NewQuery(b.db)
	err := query.UpdateBillingPoData(data)
	if err != nil {
		log.Printf("Failed to update billing PO data: %v", err)
		return err
	}
	return nil
}

func (b *BillingPoRepository) DeleteBillingPo(id int) error {
	tx, err := b.db.Begin()
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

	_, err = tx.Exec("DELETE FROM billingposubmitteddata WHERE id = $1", id)
	if err != nil {
		log.Printf("Failed to delete record with id %d: %v", id, err)
		return err
	}

	return nil
}
