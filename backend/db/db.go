package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DBHost     = os.Getenv("DB_HOST")
	DBPort     = os.Getenv("DB_PORT")
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName     = os.Getenv("DB_NAME")
)

func init() {
	//Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Check environment variables are set
	if DBHost == "" || DBPort == "" || DBUser == "" || DBPassword == "" || DBName == "" {
		log.Fatal("Database configuration environment variables are not set")
	}
}

// ConnectDB returns a new connection to the database.
func ConnectDB(initMode bool) (*sql.DB, error) {
	if !initMode {
		if err := CheckDBExists(); err != nil {
			return nil, err
		}
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)
	return sql.Open("postgres", psqlInfo)
}

// CheckDBExists checks if the database exists
func CheckDBExists() error {
	db, err := ConnectDB(true)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname=$1)"
	err = db.QueryRow(query, DBName).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("database %s does not exist", DBName)
	}

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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// GetUserID looks up the id for a username.
func GetUserID(db *sql.DB, username string) (int, error) {
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
