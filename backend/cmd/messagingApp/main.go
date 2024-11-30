package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"messagingApp/backend/pkg/api"
	"messagingApp/backend/pkg/db"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK) 
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.DB.Close() 

	router := mux.NewRouter()

	router.HandleFunc("/api/messages", api.CreateMessageHandler).Methods("POST")
	router.HandleFunc("/api/messages", api.GetPendingMessagesHandler).Methods("GET")
	router.HandleFunc("/api/messages/{id}", api.GetMessageHandler).Methods("GET")
	router.HandleFunc("/api/messages/{id}/respond", api.RespondToMessageHandler).Methods("POST")

	router.HandleFunc("/api/message/search", api.SearchMessageHandler).Methods("GET")

	corsRouter := enableCORS(router)

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", corsRouter); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
