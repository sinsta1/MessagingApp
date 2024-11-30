package models

// Agent struct represents an agent in the system
type Agent struct {
    ID        int64    `json:"id"`        // Primary key, auto-increment
    Name      string `json:"name"`      // Name of the agent
    Email     string `json:"email"`     // Email of the agent (unique)
    Phone     string `json:"phone"`     // Phone number of the agent
    CreatedAt string `json:"created_at"`// Timestamp when the agent is created
}
