package db

import (
    "messagingApp/backend/pkg/models"
    "fmt"
)

// Get all agents
func GetAllAgents() ([]models.Agent, error) {
    rows, err := DB.Query("SELECT agent_id, name FROM agents")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var agents []models.Agent
    for rows.Next() {
        var agent models.Agent
        if err := rows.Scan(&agent.ID, &agent.Name); err != nil {
            return nil, err
        }
        agents = append(agents, agent)
    }

    return agents, nil
}

// CreateAgent inserts a new agent into the database.
func CreateAgent(agent *models.Agent) error {
    query := `INSERT INTO agents (name, email, phone) VALUES (?, ?, ?)`
    result, err := DB.Exec(query, agent.Name, agent.Email, agent.Phone)
    if err != nil {
        return fmt.Errorf("failed to insert agent: %v", err)
    }

    // Retrieve the ID of the inserted agent (auto-increment).
    agent.ID, err = result.LastInsertId()
    if err != nil {
        return fmt.Errorf("failed to retrieve agent ID: %v", err)
    }

    return nil
}
