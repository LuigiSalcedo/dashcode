package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	if len(os.Args) < 3 {
		log.Fatal("Not enougth args. The program needs these args: user_name password db_name")
	}

	dbConn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", os.Args[1], os.Args[2], os.Args[3]))

	if err != nil {
		log.Fatalf("Error with database connection: %v", err)
	}

	db = dbConn
}

// Get the database connection
func DB() *sql.DB {
	return db
}

// Init a Database Transaction
func Begin() (*sql.Tx, error) {
	return db.Begin()
}
