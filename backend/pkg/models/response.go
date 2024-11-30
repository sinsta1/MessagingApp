package models

import "time"

// Response represents a response to a customer message by an agent
type Response struct {
	ID          int       `json:"id"`
	MessageID   int       `json:"message_id"`
	AgentName   string    `json:"agent_name"`
	ResponseText string   `json:"response"`
	RespondedAt time.Time `json:"responded_at"`
}
