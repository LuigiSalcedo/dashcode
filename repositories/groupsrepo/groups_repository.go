package groupsrepo

import (
	"dashcode/database"
	"dashcode/general"
	"dashcode/repositories"
	"database/sql"
)

const (
	saveSQL = "INSERT INTO GROUPS(id_creator, name, description) VALUES (?, ?, ?)"
)

var (
	saveStmt *sql.Stmt
)

func init() {
	check := func(err error) {
		if err != nil {
			repositories.Error("groups", err)
		}
	}

	stmt, err := database.DB().Prepare(saveSQL)
	saveStmt = stmt
	check(err)
}

// Save a group in the database
func CreateGroup(idCreator int64, name, description string) error {
	_, err := saveStmt.Exec(idCreator, name, description)

	if err != nil {
		general.PotencialInternalError(err)
		return err
	}

	return nil
}
