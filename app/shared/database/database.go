package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

var (
	SQL sql.DB
)

type Type string

const (
	TypeMySQL Type = "MySQL"
)

// Info contains the database configurations
type Info struct {
	// Database type
	Type Type
	// MySQL info if used
	MySQL MySQLInfo
	// Bolt info if used
}

// MySQLInfo is the details for the database connection
type MySQLInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

// Connect to the database
func Connect() *sql.DB {

	db, err := sql.Open("mysql", "ding:Pass_stfu404@/auth")
	if err != nil {
		log.Println("SQL Driver Error", err) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Println("Database Error", err) // proper error handling instead of panic in your app
	}

	fmt.Println("Database connection Stablished :))")
	return db
}
