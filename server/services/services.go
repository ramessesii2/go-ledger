package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RAMESSESII2/go-ledger/server/models"
	"github.com/RAMESSESII2/go-ledger/server/repositories"
	"github.com/gorilla/mux"
)

func GetLedger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "hi!")
	var ledger []models.Transaction
	repositories.DB.Find(&ledger)
	json.NewEncoder(w).Encode(ledger)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi!")
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var trasaction models.Transaction
	repositories.DB.First(&trasaction, params["id"])
	if trasaction.ID == 0 {
		return
	}
	json.NewEncoder(w).Encode(trasaction)
}
func NewTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var trasaction models.Transaction
	json.NewDecoder(r.Body).Decode(&trasaction)
	if trasaction.FirstName == "" {
		return
	}
	repositories.DB.Create(&trasaction)
	json.NewEncoder(w).Encode(trasaction)
}
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var trasaction models.Transaction
	repositories.DB.First(&trasaction, params["id"])
	if trasaction.ID <= 0 {
		return
	}
	var newTransaction models.Transaction
	json.NewDecoder(r.Body).Decode(&newTransaction)
	newTransaction.ID = trasaction.ID
	if newTransaction.FirstName == "" {
		newTransaction.FirstName = trasaction.FirstName
	}
	if newTransaction.LastName == "" {
		newTransaction.LastName = trasaction.LastName
	}
	repositories.DB.Save(&newTransaction)
	json.NewEncoder(w).Encode(newTransaction)
}
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var trasaction models.Transaction
	repositories.DB.Delete(&trasaction, params["id"])
	json.NewEncoder(w).Encode("The transaction is deleted successfully!")
}
