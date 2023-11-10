package usersrepo

import (
	"dashcode/database"
	"dashcode/general"
	"dashcode/repositories"
	"dashcode/repositories/loginrepo"
	"database/sql"
)

// Repository name
const repoName = "users"

// SQL Queries
const (
	saveSQL = "INSERT INTO USERS (ID, NAME) VALUES(?, ?)"
)

// Prepared Statement
var (
	saveStmt *sql.Stmt
)

func init() {
	check := func(err error) {
		if err != nil {
			repositories.Error(repoName, err)
		}
	}
	stmt, err := database.DB().Prepare(saveSQL)
	saveStmt = stmt
	check(err)
}

// Save an user in User and Login
func SaveUser(id int64, name, email string, password []byte) error {
	tx, err := database.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	if err != nil {
		return err
	}

	stmt := tx.Stmt(saveStmt)

	_, err = stmt.Exec(id, name)

	if err != nil {
		general.PotencialInternalError(err)
		return err
	}

	err = loginrepo.SaveLoginWithTx(tx, id, email, []byte(password))

	if err != nil {
		general.PotencialInternalError(err)
		return err
	}

	return nil
}
