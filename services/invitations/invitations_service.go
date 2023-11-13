package invitations

import (
	"dashcode/database"
	"dashcode/general"
	"dashcode/models"
	"dashcode/repositories/groupsrepo"
	"dashcode/repositories/invitationsrepo"
	"dashcode/services"
	"net/http"
)

type InvitationState struct {
	State bool
}

var (
	Accepted = &InvitationState{State: true}
	Rejected = &InvitationState{State: false}
)

// Special responses
var (
	AlreadyResponded = &services.SpecialResponse{
		Code:    http.StatusBadRequest,
		Message: "invitations already responded",
	}
)

func FetchInvitations(id int64) ([]models.InvitationData, *services.Error) {
	r, err := invitationsrepo.FetchById(id)

	if err != nil {
		return nil, services.ErrorInternal
	}

	return r, nil
}

func FetchGroupInvitations(ownerId, groupId int64) ([]models.SentInvitationsData, *services.Error) {
	id, err := groupsrepo.FetchOwner(groupId)

	if err != nil {
		return nil, services.ErrorInternal
	}

	if id != ownerId {
		return nil, services.ErrorForbidden
	}

	invs, err := invitationsrepo.FetchAllByGroupId(groupId)

	if err != nil {
		return nil, services.ErrorInternal
	}

	return invs, nil
}

func FetchInvitationsWithState(ownerId, groupId int64, state *InvitationState) ([]models.SentInvitationsData, *services.Error) {
	if state == nil {
		return FetchGroupInvitations(ownerId, groupId)
	}

	id, err := groupsrepo.FetchOwner(groupId)

	if err != nil {
		return nil, services.ErrorInternal
	}

	if id != ownerId {
		return nil, services.ErrorForbidden
	}

	var invs []models.SentInvitationsData

	if state == nil {
		invs, err = invitationsrepo.FetchNullByGroup(groupId)
	} else {
		invs, err = invitationsrepo.FetchWithStateByGroup(groupId, state.State)
	}

	if err != nil {
		return nil, services.ErrorInternal
	}

	return invs, nil
}

func FetchUserInvitations(userId int64) ([]models.UserInvitationData, *services.Error) {
	r, err := invitationsrepo.FetchUserInvitations(userId)

	if err != nil {
		return nil, services.ErrorInternal
	}

	return r, nil
}

func RespondInvitation(userId, invitationId int64, response bool) (*services.SpecialResponse, *services.Error) {
	id, groupId, responded, err := invitationsrepo.FetchInvitation(invitationId)

	if err != nil {
		return nil, services.ErrorInternal
	}

	if id != userId {
		return nil, services.ErrorForbidden
	}

	if responded {
		return AlreadyResponded, nil
	}

	tx, err := database.Begin()

	if err != nil {
		general.PotencialInternalError(err)
		return nil, services.ErrorInternal
	}

	err = invitationsrepo.UpdateState(tx, response, invitationId)

	if err != nil {
		tx.Rollback()
		return nil, services.ErrorInternal
	}

	if !response {
		tx.Commit()
		return nil, nil
	}

	err = invitationsrepo.SaveGroupMember(tx, groupId, userId)

	if err != nil {
		tx.Rollback()
		return nil, services.ErrorInternal
	}

	tx.Commit()
	return nil, nil
}
