package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Naveeshkumar24/internal/models"
	"github.com/Naveeshkumar24/pkg/utils"
	"github.com/Naveeshkumar24/repository"
	"github.com/gorilla/mux"
)

type BillingPoHandler struct {
	billingRepo *repository.BillingPoRepository
}

func NewBillingPoHandler(billingRepo models.BillingPoInterface) *BillingPoHandler {
	return &BillingPoHandler{
		billingRepo: billingRepo.(*repository.BillingPoRepository),
	}
}

func (b *BillingPoHandler) FetchDropDown(w http.ResponseWriter, r *http.Request) {
	dropdownList, err := b.billingRepo.FetchDropDown()
	if err != nil {
		log.Printf("Error fetching dropdown data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.Encode(w, map[string]string{"message": "Internal server error"})
		return
	}
	if len(dropdownList) == 0 {
		log.Printf("No dropdown data found")
		w.WriteHeader(http.StatusNotFound)
		utils.Encode(w, map[string]string{"message": "No data found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	utils.Encode(w, dropdownList)
}

func (b *BillingPoHandler) SubmitFormBillingPoData(w http.ResponseWriter, r *http.Request) {
	var data models.BillingPo
	err := utils.Decode(r, &data)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		utils.Encode(w, map[string]string{"message": "Invalid request body"})
		return
	}

	err = b.billingRepo.SubmitFormBillingPoData(data)
	if err != nil {
		log.Printf("Failed to submit billing PO data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.Encode(w, map[string]string{"message": "Failed to submit data"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.Encode(w, map[string]string{"message": "Data submitted successfully"})
}

func (b *BillingPoHandler) FetchBillingPoData(w http.ResponseWriter, r *http.Request) {
	billingPoList, err := b.billingRepo.FetchBillingPoData(r)
	if err != nil {
		log.Printf("Failed to fetch billing PO data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.Encode(w, map[string]string{"message": "Internal server error"})
		return
	}
	if len(billingPoList) == 0 {
		log.Printf("No data found")
		w.WriteHeader(http.StatusNotFound)
		utils.Encode(w, map[string]string{"message": "No data found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	utils.Encode(w, billingPoList)
}

func (b *BillingPoHandler) UpdateBillingPoData(w http.ResponseWriter, r *http.Request) {
	var data models.BillingPo
	err := utils.Decode(r, &data)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		utils.Encode(w, map[string]string{"message": "Invalid request body"})
		return
	}
	err = b.billingRepo.UpdateBillingPoData(data)
	if err != nil {
		log.Printf("Failed to update billing PO data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.Encode(w, map[string]string{"message": "Failed to update data"})
		return
	}
	w.WriteHeader(http.StatusOK)
	utils.Encode(w, map[string]string{"message": "Data updated successfully"})
}

func (b *BillingPoHandler) DeleteBillingPoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = b.billingRepo.DeleteBillingPoData(id)
	if err != nil {
		log.Printf("Error deleting record: %v", err)
		http.Error(w, "Failed to delete record", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Record deleted successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
