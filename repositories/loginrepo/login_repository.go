package loginrepo

import (
	"dashcode/database"
	"dashcode/general"
	"dashcode/repositories"
	"database/sql"
)

// Get the ID of an user using the login data
var (
	fetchId   *sql.Stmt
	saveLogin *sql.Stmt
)

// Repository name
var repoName = "login"

// SQL Queries
const (
	fetchIdSQL = "SELECT ID, PASSWORD FROM LOGIN WHERE EMAIL = ?"
	saveSQL    = "INSERT INTO LOGIN(ID, EMAIL, PASSWORD) VALUES(?, ?, ?)"
)

// Check repository error
func checkError(err error) {
	if err != nil {
		repositories.Error(repoName, err)
	}
}

func init() {
	stmt, err := database.DB().Prepare(fetchIdSQL)
	checkError(err)
	fetchId = stmt

	stmt, err = database.DB().Prepare(saveSQL)
	checkError(err)
	saveLogin = stmt
}

// Save a Login in database using database transaction
func SaveLoginWithTx(tx *sql.Tx, id int64, email string, password []byte) error {
	stmt := tx.Stmt(saveLogin)

	_, err := stmt.Exec(id, email, password)

	if err != nil {
		general.PotencialInternalError(err)
	}

	return err
}

// Save a login in database
func SaveLogin(id int64, email string, password []byte) error {
	tx, err := database.Begin()

	if err != nil {
		general.PotencialInternalError(err)
		return err
	}

	return SaveLoginWithTx(tx, id, email, password)
}

// Get ID and Password from Login using Email
func FetchLogin(email string) (int64, []byte, error) {
	var id int64
	var hash []byte

	r, err := fetchId.Query(email)

	if err != nil {
		return -1, nil, err
	}

	if !r.Next() {
		return -1, nil, nil
	}

	r.Scan(&id, &hash)

	return id, hash, nil
}
