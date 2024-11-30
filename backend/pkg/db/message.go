package db

import (
	"database/sql"
	"fmt"
	"messagingApp/backend/pkg/models"
	"time"
)

// Create a new customer message
func CreateMessage(m *models.Message) (int, error) {
    query := `
        INSERT INTO messages (customer_name, customer_email, message, status, created_at)
        VALUES (?, ?, ?, ?, ?)
    `
    result, err := DB.Exec(query, m.CustomerName, m.CustomerEmail, m.MessageText, "pending", time.Now())
    if err != nil {
        return 0, err
    }

    // Retrieve the ID of the newly inserted message
    messageID, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(messageID), nil
}

func GetPendingMessages() ([]models.Message, error) {
	rows, err := DB.Query(`
		SELECT id, customer_name, customer_email, message, status
		FROM messages
		WHERE status = 'pending'
		ORDER BY created_at ASC
	`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(
			&message.ID,
			&message.CustomerName,
			&message.CustomerEmail,
			&message.MessageText,
			&message.Status,
			// &message.CreatedAt,
		); err != nil {
			fmt.Println(err)
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func GetMessageByID(id int) (*models.Message, []models.Response, error) {
	// Query the message
	message := &models.Message{}
	err := DB.QueryRow(`
		SELECT id, customer_name, customer_email, message, status
		FROM messages
		WHERE id = ?`, id).Scan(
		&message.ID,
		&message.CustomerName,
		&message.CustomerEmail,
		&message.MessageText,
		&message.Status,
		// &message.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, nil // No message found
		}
		return nil, nil, err
	}

	// Query the responses for the message
	rows, err := DB.Query(`
		SELECT id, message_id, agent_name, response, responded_at
		FROM responses
		WHERE message_id = ?`, id)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var responses []models.Response
	for rows.Next() {
		var response models.Response
		if err := rows.Scan(
			&response.ID,
			&response.MessageID,
			&response.AgentName,
			&response.ResponseText,
			&response.RespondedAt,
		); err != nil {
			return nil, nil, err
		}
		responses = append(responses, response)
	}

	return message, responses, nil
}

