package groupsrepo

import (
	"dashcode/database"
	"dashcode/general"
	"dashcode/models"
	"dashcode/repositories"
	"dashcode/repositories/loginrepo"
	"database/sql"
)

const (
	saveSQL           = "INSERT INTO GROUPS(id_creator, name, description) VALUES (?, ?, ?)"
	saveInvitationSQL = "INSERT INTO INVITATIONS(id_user, id_group) VALUES(?, ?)"
	fetchByOwnerSQL   = "SELECT ID, ID_CREATOR, NAME, DESCRIPTION FROM GROUPS WHERE ID_CREATOR = ?"
	fetchByMemberSQL  = `
	SELECT G.ID, G.ID_CREATOR, G.NAME, G.DESCRIPTION
	FROM GROUPS as G JOIN GROUP_MEMBERS as GM ON GM.ID_GROUP = G.ID
	WHERE GM.ID_USER = ?
	`
	fetchOwnerSQL       = "SELECT ID_CREATOR FROM GROUPS WHERE ID = ? LIMIT 1"
	fetchInvitationsSQL = "SELECT STATE FROM INVITATIONS WHERE ID_USER = ? AND ID_GROUP = ?"
)

var (
	saveStmt             *sql.Stmt
	saveInvitationStmt   *sql.Stmt
	fetchByOwnerStmt     *sql.Stmt
	fetchByMemberStmt    *sql.Stmt
	fetchOwnerStmt       *sql.Stmt
	fetchInvitationsStmt *sql.Stmt
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

	stmt, err = database.DB().Prepare(saveInvitationSQL)
	saveInvitationStmt = stmt
	check(err)

	stmt, err = database.DB().Prepare(fetchByOwnerSQL)
	fetchByOwnerStmt = stmt
	check(err)

	stmt, err = database.DB().Prepare(fetchByMemberSQL)
	fetchByMemberStmt = stmt
	check(err)

	stmt, err = database.DB().Prepare(fetchOwnerSQL)
	fetchOwnerStmt = stmt
	check(err)

	stmt, err = database.DB().Prepare(fetchInvitationsSQL)
	fetchInvitationsStmt = stmt
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

func FetchByOwner(idCreator int64) ([]models.Group, error) {
	r, err := fetchByOwnerStmt.Query(idCreator)

	if err != nil {
		general.PotencialInternalError(err)
		return nil, err
	}

	result := make([]models.Group, 0, 15)

	for r.Next() {
		g := models.Group{}
		err = r.Scan(&g.Id, &g.IdCreator, &g.Name, &g.Description)
		result = append(result, g)

		if err != nil {
			general.PotencialInternalError(err)
			return nil, err
		}
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

func FetchByMember(idMember int64) ([]models.Group, error) {
	r, err := fetchByMemberStmt.Query(idMember)

	if err != nil {
		general.PotencialInternalError(err)
		return nil, err
	}

	groups := make([]models.Group, 0, 15)

	for r.Next() {
		g := models.Group{}
		err = r.Scan(&g.Id, &g.IdCreator, &g.Name, &g.Description)

		if err != nil {
			general.PotencialInternalError(err)
			return nil, err
		}

		groups = append(groups, g)
	}

	if len(groups) == 0 {
		return nil, nil
	}

	return groups, nil
}

func FetchOwner(idGroup int64) (int64, error) {
	r, err := fetchOwnerStmt.Query(idGroup)

	if err != nil {
		general.PotencialInternalError(err)
		return -1, err
	}

	if !r.Next() {
		return -1, nil
	}

	var idCreator int64

	err = r.Scan(&idCreator)

	if err != nil {
		return -1, err
	}

	return idCreator, nil
}

func FetchCurrentInvitation(email string, idGroup int64) (int64, *sql.NullBool, error) {
	idUser, _, err := loginrepo.FetchLogin(email)

	if idUser == -1 {
		return -1, nil, err
	}

	r, err := fetchInvitationsStmt.Query(idUser, idGroup)

	if err != nil {
		general.PotencialInternalError(err)
		return -1, nil, err
	}

	state := sql.NullBool{}

	if !r.Next() {
		return idUser, nil, nil
	}

	err = r.Scan(&state)

	if err != nil {
		general.PotencialInternalError(err)
		return -1, nil, err
	}

	return idUser, &state, nil
}

func SendInvitation(idUser, idGroup int64) error {
	_, err := saveInvitationStmt.Exec(idUser, idGroup)

	if err != nil {
		general.PotencialInternalError(err)
		return err
	}

	return nil
}
