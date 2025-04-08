package handlers

import (
	"encoding/json" // For JSON encoding/decoding
	"fmt"
	"net/http" // Standard HTTP interfaces
	"strconv"

	"github.com/gorilla/mux" // The router being used
	"github.com/yagyansh5/simplebank/db"
	"github.com/yagyansh5/simplebank/models"
)

// store holds the database connection/querier.
// Ensure this is initialized appropriately elsewhere in your application.

// TransferTxHandler handles initiating a transfer transaction.
// It expects JSON body with db.TransferTxParams fields.

// CreateAccountHandler creates a new account.
// It expects a JSON body with Owner, Balance, and Currency.
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Owner    string `json:"owner"`
		Balance  int64  `json:"balance"`
		Currency string `json:"currency"`
	}

	// Decode the request body into the request struct
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input: %v", err), http.StatusBadRequest)
		return
	}

	// Validate input (basic example)
	if req.Owner == "" || req.Currency == "" {
		http.Error(w, "Owner and Currency fields are required", http.StatusBadRequest)
		return
	}

	// Create the account in the database
	newAccount, err := db.CreateAccount(db.DB, models.Account{
		Owner:    req.Owner,
		Balance:  req.Balance,
		Currency: req.Currency,
	})
	if err != nil {
		// Handle potential DB errors (e.g., unique constraint violation)
		http.Error(w, fmt.Sprintf("Error creating account: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the created account
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	if err := json.NewEncoder(w).Encode(newAccount); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
	}
}

// GetAccountHandler fetches an account by ID from the URL path.
func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountIDStr := vars["id"] // Assumes route is defined like "/accounts/{id}"

	// Convert the account ID from string to int64
	id, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid account ID format", http.StatusBadRequest)
		return
	}

	// Fetch the account from the database
	account, err := db.GetAccount(db.DB, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching account: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the account details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	if err := json.NewEncoder(w).Encode(account); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
	}
}

// CreateEntryHandler creates a new entry (e.g., a deposit or withdrawal).
// It expects a JSON body with AccountID and Amount.
func CreateEntryHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AccountID int64 `json:"account_id"`
		Amount    int64 `json:"amount"`
	}

	// Decode the request body
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input: %v", err), http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.AccountID <= 0 {
		http.Error(w, "Invalid Account ID", http.StatusBadRequest)
		return
	}

	// Create the entry in the database
	newEntry, err := db.CreateEntry(db.DB, models.Entry{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	})
	if err != nil {
		// Handle DB errors (e.g., foreign key constraint if account doesn't exist)
		http.Error(w, fmt.Sprintf("Error creating entry: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the created entry
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	if err := json.NewEncoder(w).Encode(newEntry); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
	}
}

// GetEntryHandler fetches an entry by ID from the URL path.
func GetEntryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryIDStr := vars["id"] // Assumes route is defined like "/entries/{id}"

	// Convert the entry ID from string to int64
	id, err := strconv.ParseInt(entryIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid entry ID format", http.StatusBadRequest)
		return
	}

	// Fetch the entry from the database
	entry, err := db.GetEntry(db.DB, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching entry: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the entry details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	if err := json.NewEncoder(w).Encode(entry); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
	}
}

// CreateTransferHandler handles creating a transfer between accounts using the TransferTx database method.
// It expects a JSON body with db.TransferTxParams fields.
