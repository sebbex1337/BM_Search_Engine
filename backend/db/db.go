package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var DatabasePath string

func init() {
	//Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the database path from the environment variable
	DatabasePath = os.Getenv("DATABASE_PATH")
	if DatabasePath == "" {
		log.Fatal("DATABASE_PATH environment variable is not set")
	}
}

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

	migration, err := os.ReadFile("./migration.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(migration))
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

	// Initialize an empty slice to hold the query results
	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		// Creates a slice to hold the column values
		values := make([]interface{}, len(columns))
		// Creates a slice of pointers to the column values
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		// Creates a map to hold the column name-value pairs
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
