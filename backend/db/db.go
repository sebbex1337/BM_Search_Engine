package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DatabasePath = "./whoknows.db"

// ConnectDB returns a new connection to the database.
func ConnectDB(initMode bool) (*sql.DB, error) {
    if !initMode {
        if err := CheckDBExists(); err != nil {
            return nil, err
        }
    }
    return sql.Open("sqlite3", DatabasePath)
}

// CheckDBExists checks if the database file exists.
func CheckDBExists() error {
    if _, err := os.Stat(DatabasePath); os.IsNotExist(err) {
        return fmt.Errorf("database not found")
    }
    return nil
}

// InitDB initializes the database tables.
func InitDB() error {
    db, err := ConnectDB(true)
    if err != nil {
        return err
    }
    defer db.Close()

    schema, err := os.ReadFile("./schema.sql")
    if err != nil {
        return err
    }

    _, err = db.Exec(string(schema))
    if err != nil {
        return err
    }

    log.Println("Initialized the database:", DatabasePath)
    return nil
}

// QueryDB queries the database and returns a list of maps.
func QueryDB(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, error) {
    rows, err := db.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    columns, err := rows.Columns()
    if err != nil {
        return nil, err
    }

    results := make([]map[string]interface{}, 0)
    for rows.Next() {
        values := make([]interface{}, len(columns))
        valuePtrs := make([]interface{}, len(columns))
        for i := range values {
            valuePtrs[i] = &values[i]
        }

        if err := rows.Scan(valuePtrs...); err != nil {
            return nil, err
        }

        result := make(map[string]interface{})
        for i, col := range columns {
            result[col] = values[i]
        }
        results = append(results, result)
    }

    return results, nil
}

// GetUserID looks up the id for a username.
func GetUserID(db *sql.DB, username string) (int, error) {
    var id int
    err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}