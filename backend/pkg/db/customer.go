package db

import (
    "messagingApp/backend/pkg/models"
    "fmt"
)

// Get all customers
func GetAllCustomers() ([]models.Customer, error) {
    rows, err := DB.Query("SELECTid, name, email FROM customers")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var customers []models.Customer
    for rows.Next() {
        var customer models.Customer
        if err := rows.Scan(&customer.ID, &customer.Name, &customer.Email); err != nil {
            return nil, err
        }
        customers = append(customers, customer)
    }

    return customers, nil
}

// CreateCustomer inserts a new customer into the database
func CreateCustomer(customer *models.Customer) error {
    query := "INSERT INTO customers (name, email, phone) VALUES (?, ?, ?)"
    result, err := DB.Exec(query, customer.Name, customer.Email, customer.Phone)
    if err != nil {
        return fmt.Errorf("failed to create customer: %v", err)
    }

    // Retrieve the last inserted ID and assign it to the customer struct
    id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("failed to retrieve last insert ID: %v", err)
    }

    customer.ID = int(id) // Update customer ID with the auto-generated value
    return nil
}

