package models

import "time"

// Message represents a customer message in the system
type Message struct {
    ID            int       `json:"id"`
    CustomerName  string    `json:"customer_name"`
    CustomerEmail string    `json:"customer_email,omitempty"`
    MessageText   string    `json:"message"`
    Status        string    `json:"status"`
    CreatedAt     time.Time `json:"created_at"`
}
