package api

import (
    "encoding/json"
    "net/http"
    "messagingApp/backend/pkg/db"
    "messagingApp/backend/pkg/models"
)

// HandleAgents is a handler function for creating new agents and fetching all agents.
func HandleAgents(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        agents, err := db.GetAllAgents() // Get all agents
        if err != nil {
            http.Error(w, "Failed to fetch agents", http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(agents) // Return all agents in response
    } else if r.Method == http.MethodPost {
        var agent models.Agent // Declare a new agent instance
        err := json.NewDecoder(r.Body).Decode(&agent) // Decode request body into agent struct
        if err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        err = db.CreateAgent(&agent) // Call the database function to insert the agent
        if err != nil {
            http.Error(w, "Failed to create agent", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated) // Send back a 201 response code
        json.NewEncoder(w).Encode(agent) // Return the created agent in response
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}
