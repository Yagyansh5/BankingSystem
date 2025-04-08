package main

import (
	"log"
	"net/http"

	"github.com/yagyansh5/simplebank/db"
	"github.com/yagyansh5/simplebank/routes"
	"github.com/yagyansh5/simplebank/utils"

	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables and initialize services
	utils.LoadEnv()
	utils.InitLogger()
	db.InitDB()

	// Initialize Mux router
	router := mux.NewRouter()

	router.Use(CORSMiddleware)

	// Set up routes
	routes.SetupRouter(router)

	// Set up CORS headers manually as Mux doesn't have middleware like Fiber
	// Create a handler that allows cross-origin requests
	router.Use(mux.CORSMethodMiddleware(router))

	// Start the HTTP server
	utils.Logger.Info("Starting User Service on port 3001")
	log.Fatal(http.ListenAndServe(":3001", router))
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
