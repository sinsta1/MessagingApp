// db/response.go
package db

import (
    "time"
)

// CreateResponse stores a response for a specific message.
func SaveResponse(messageID int, agentName, response string) error {
    query := `INSERT INTO responses (message_id, agent_name, response, responded_at) VALUES (?, ?, ?, ?)`
    _, err := DB.Exec(query, messageID, agentName, response, time.Now())
    return err
}


// UpdateMessageStatus updates the status of a message to 'responded'.
func UpdateMessageStatus(messageID int, status string) error {
    query := `UPDATE messages SET status = ? WHERE id = ?`
    _, err := DB.Exec(query, status, messageID)
    return err
}
