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

	fetchAllByGroupSQL = ` 
	SELECT I.ID, U.NAME, L.EMAIL, I.STATE
	FROM
	INVITATIONS as I 
	JOIN GROUPS as G ON I.ID_GROUP = G.ID
	JOIN USERS as U ON I.ID_USER = U.ID
	JOIN LOGIN as L ON L.ID = U.ID
	WHERE G.ID = ?
	`

	fetchWithStateByGroupSQL = `
	SELECT I.ID, U.NAME, L.EMAIL, I.STATE
	FROM
	INVITATIONS as I 
	JOIN GROUPS as G ON I.ID_GROUP = G.ID
	JOIN USERS as U ON I.ID_USER = U.ID
	JOIN LOGIN as L ON L.ID = U.ID
	WHERE G.ID = ? AND STATE = ?	
	`

	fetchNullByGroupSQL = `
	SELECT I.ID, U.NAME, L.EMAIL, I.STATE
	FROM
	INVITATIONS as I 
	JOIN GROUPS as G ON I.ID_GROUP = G.ID
	JOIN USERS as U ON I.ID_USER = U.ID
	JOIN LOGIN as L ON L.ID = U.ID
	WHERE G.ID = ? AND STATE IS NULL
	`

	fetchUserInvitationsSQL = `
	SELECT 
	I.ID,
	G.ID,
	G.NAME,
	G.DESCRIPTION
	FROM
	INVITATIONS as I
	JOIN GROUPS as G ON I.ID_GROUP = G.ID
	JOIN USERS as U ON I.ID_USER = U.ID
	WHERE U.ID = ? AND I.STATE IS NULL
	`

	fetchInvitadedSQL = `
	SELECT
	ID_USER, ID_GROUP, STATE
	FROM INVITATIONS
	WHERE ID = ? LIMIT 1
	`

	updateStateSQL = `
	UPDATE INVITATIONS SET STATE = ? WHERE ID = ?
	`

	saveGroupMemberSQL = `
	INSERT INTO
	GROUP_MEMBERS (id_group, id_user)
	VALUES (?, ?)
	`
)

var (
	fetchByIdStmt             *sql.Stmt
	fetchAllByGroupIdStmt     *sql.Stmt
	fetchWithStateByGroupStmt *sql.Stmt
	fetchNullByGroupStmt      *sql.Stmt
	fetchUserInvitationsStmt  *sql.Stmt
	fetchInvitatedStmt        *sql.Stmt
	updateStateStmt           *sql.Stmt
	saveGroupMemberStmt       *sql.Stmt
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

	stmt, err = database.DB().Prepare(fetchAllByGroupSQL)
	fetchAllByGroupIdStmt = stmt
	check(err)

	stmt, err = database.DB().Prepare(fetchWithStateByGroupSQL)
	fetchWithStateByGroupStmt = stmt
	check(err)

	stmt, err = database.DB().Prepare(fetchNullByGroupSQL)
	fetchNullByGroupStmt = stmt
	check(err)

	stmt, err = database.DB().Prepare(fetchUserInvitationsSQL)
	fetchUserInvitationsStmt = stmt
	check(err)

	stmt, err = database.DB().Prepare(fetchInvitadedSQL)
	fetchInvitatedStmt = stmt
	check(err)

	stmt, err = database.DB().Prepare(updateStateSQL)
	updateStateStmt = stmt
	check(err)

	stmt, err = database.DB().Prepare(saveGroupMemberSQL)
	saveGroupMemberStmt = stmt
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

func fetchByGroup(queryStmt *sql.Stmt, queryArgs ...any) ([]models.SentInvitationsData, error) {
	r, err := queryStmt.Query(queryArgs...)

	if err != nil {
		general.PotencialInternalError(err)
		return nil, err
	}

	invs := make([]models.SentInvitationsData, 0, 15)

	for r.Next() {
		inv := models.SentInvitationsData{}
		state := sql.NullBool{}
		err = r.Scan(&inv.Id, &inv.UserName, &inv.UserEmail, &state)

		if err != nil {
			general.PotencialInternalError(err)
			return nil, err
		}

		if !state.Valid {
			inv.State = "No response"
			invs = append(invs, inv)
			continue
		}

		if state.Bool {
			inv.State = "Accepted"
		} else {
			inv.State = "Rejected"
		}

		invs = append(invs, inv)
	}

	return invs, nil
}

func FetchAllByGroupId(groupId int64) ([]models.SentInvitationsData, error) {
	return fetchByGroup(fetchAllByGroupIdStmt, groupId)
}

func FetchWithStateByGroup(groupId int64, state bool) ([]models.SentInvitationsData, error) {
	return fetchByGroup(fetchWithStateByGroupStmt, groupId, state)
}

func FetchNullByGroup(groupId int64) ([]models.SentInvitationsData, error) {
	return fetchByGroup(fetchNullByGroupStmt, groupId)
}

func FetchUserInvitations(userId int64) ([]models.UserInvitationData, error) {
	r, err := fetchUserInvitationsStmt.Query(userId)

	if err != nil {
		general.PotencialInternalError(err)
		return nil, err
	}

	invs := make([]models.UserInvitationData, 0, 15)

	for r.Next() {
		inv := models.UserInvitationData{}

		err = r.Scan(&inv.Id, &inv.GroupId, &inv.GroupName, &inv.GroupDescription)

		if err != nil {
			general.PotencialInternalError(err)
			return nil, err
		}

		invs = append(invs, inv)
	}

	return invs, nil
}

func FetchInvitation(invitationId int64) (int64, int64, bool, error) {
	r, err := fetchInvitatedStmt.Query(invitationId)

	if err != nil {
		general.PotencialInternalError(err)
		return -1, -1, false, err
	}

	var userId int64
	var groupId int64
	nullableState := sql.NullBool{}

	if !r.Next() {
		return -1, -1, false, nil
	}

	err = r.Scan(&userId, &groupId, &nullableState)

	if err != nil {
		general.PotencialInternalError(err)
		return -1, -1, false, nil
	}

	responded := nullableState.Valid

	return userId, groupId, responded, nil
}

func UpdateState(tx *sql.Tx, state bool, invitationId int64) error {
	stmt := tx.Stmt(updateStateStmt)

	_, err := stmt.Exec(state, invitationId)

	if err != nil {
		general.PotencialInternalError(err)
	}

	return err
}

func SaveGroupMember(tx *sql.Tx, groupId, userId int64) error {
	stmt := tx.Stmt(saveGroupMemberStmt)

	_, err := stmt.Exec(groupId, userId)

	if err != nil {
		general.PotencialInternalError(err)
	}

	return err
}
