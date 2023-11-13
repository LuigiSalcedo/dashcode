package invitations

import (
	"dashcode/models"
	"dashcode/repositories/groupsrepo"
	"dashcode/repositories/invitationsrepo"
	"dashcode/services"
)

type InvitationState struct {
	State bool
}

var (
	Accepted = &InvitationState{State: true}
	Rejected = &InvitationState{State: false}
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
