package routes

import (
	"github.com/gorilla/mux"
	"github.com/yagyansh5/simplebank/handlers"
)

// SetupRouter sets up the router and routes for all entities
func SetupRouter(router *mux.Router) {

	// Account routes
	router.HandleFunc("/account", handlers.CreateAccountHandler).Methods("POST")
	router.HandleFunc("/account/{id:[0-9]+}", handlers.GetAccountHandler).Methods("GET")

	// Entry routes
	router.HandleFunc("/entry", handlers.CreateEntryHandler).Methods("POST")
	router.HandleFunc("/entry/{id:[0-9]+}", handlers.GetEntryHandler).Methods("GET")

	// Transfer routes
	//router.HandleFunc("/transfer", handlers.CreateTransferHandler).Methods("POST")

}
