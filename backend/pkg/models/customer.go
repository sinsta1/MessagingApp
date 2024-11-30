package models

// Customer struct represents a customer in the system
type Customer struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Phone     string `json:"phone"`
    CreatedAt string `json:"created_at"` // Can be time.Time depending on the format you prefer
}
