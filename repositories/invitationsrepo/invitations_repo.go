package invitationsrepo

import (
	"dashcode/database"
	"dashcode/general"
	"dashcode/models"
	"dashcode/repositories"
	"database/sql"
)

const (
	fetchByIdSQL = `
	SELECT
	I.ID,
	G.ID, G.NAME, G.DESCRIPTION,
	U.NAME 
	FROM
	INVITATIONS as I JOIN GROUPS as G ON I.ID_GROUP = G.ID
	JOIN USERS as U ON U.ID = G.ID_CREATOR
	WHERE I.ID_USER = ? AND I.STATE IS NULL
	`
)

var (
	fetchByIdStmt *sql.Stmt
)

func init() {
	check := func(err error) {
		if err != nil {
			repositories.Error("invitations", err)
		}
	}

	stmt, err := database.DB().Prepare(fetchByIdSQL)
	fetchByIdStmt = stmt
	check(err)

}

func FetchById(id int64) ([]models.InvitationData, error) {
	r, err := fetchByIdStmt.Query(id)

	if err != nil {
		general.PotencialInternalError(err)
		return nil, err
	}

	invs := make([]models.InvitationData, 0, 15)

	for r.Next() {
		inv := models.InvitationData{}

		err = r.Scan(&inv.Id, &inv.GroupId, &inv.GroupName, &inv.GroupDescription, &inv.CreatorName)

		if err != nil {
			general.PotencialInternalError(err)
			return nil, err
		}

		invs = append(invs, inv)
	}

	return invs, nil
}
