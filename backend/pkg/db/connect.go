package db

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect function to initialize database connection
func Connect() error {
    // Hard-coded database credentials (not recommended)
    dbUser := "root"
    dbPassword := "waheguru"
    dbHost := "127.0.0.1"
    dbPort := "3306"
    dbName := "messagingapp"

    // Build the connection string
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    // Open the database connection
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("failed to connect to database: %v", err)
    }

    // Ping the database to check if the connection is successful
    err = DB.Ping()
    if err != nil {
        return fmt.Errorf("failed to ping the database: %v", err)
    }

    log.Println("Connected to the database successfully!")
    return nil
}
