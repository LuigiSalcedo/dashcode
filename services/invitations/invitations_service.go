package invitations

import (
	"dashcode/models"
	"dashcode/repositories/invitationsrepo"
	"dashcode/services"
)

func FetchInvitations(id int64) ([]models.InvitationData, *services.Error) {
	r, err := invitationsrepo.FetchById(id)

	if err != nil {
		return nil, services.ErrorInternal
	}

	return r, nil
}
