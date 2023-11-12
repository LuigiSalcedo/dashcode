package groups

import (
	"dashcode/models"
	"dashcode/repositories/groupsrepo"
	"dashcode/services"
	"database/sql"
	"net/http"
)

// Special responses
var (
	SpecialNoComplete = services.SpecialResponse{
		Code:    http.StatusOK,
		Message: "some invitations were not sent",
	}

	SpecialNothing = services.SpecialResponse{
		Code:    http.StatusOK,
		Message: "nobody was invited",
		Data:    "empty email list",
	}
)

func CreateGroup(idCreator int64, name, description string) *services.Error {
	err := groupsrepo.CreateGroup(idCreator, name, description)

	if err != nil {
		return services.ErrorBadRequest
	}

	return nil
}

func FetchGroupsByOwner(idCreator int64) ([]models.Group, *services.Error) {
	group, err := groupsrepo.FetchByOwner(idCreator)

	if err != nil {
		return nil, services.ErrorInternal
	}

	if group == nil {
		return nil, services.ErrorNotFound
	}

	return group, nil
}

func FetchByMember(idMember int64) ([]models.Group, *services.Error) {
	groups, err := groupsrepo.FetchByMember(idMember)

	if err != nil {
		return nil, services.ErrorInternal
	}

	if groups == nil {
		return nil, services.ErrorNotFound
	}

	return groups, nil
}

func Invite(idSender int64, inv models.Invitation) (*services.SpecialResponse, *services.Error) {
	if len(inv.Emails) == 0 {
		return &SpecialNothing, nil
	}

	idOwner, err := groupsrepo.FetchOwner(inv.IdGroup)

	if err != nil {
		return nil, services.ErrorInternal
	}

	if idOwner != idSender {
		return nil, services.ErrorUnauthorized
	}

	response := services.SpecialResponse{
		Code:    SpecialNoComplete.Code,
		Message: SpecialNoComplete.Message,
	}

	response.Data = make([]string, 0, len(inv.Emails))

	for _, email := range inv.Emails {
		idUser, state, err := groupsrepo.FetchCurrentInvitation(email, inv.IdGroup)

		if err != nil {
			return nil, services.ErrorInternal
		}

		if idUser == -1 {
			response.Data = append(response.Data.([]string), email)
			continue
		}

		if state == nil {
			state = &sql.NullBool{Valid: true, Bool: false}
		}

		if !state.Valid || (state.Valid && state.Bool) {
			response.Data = append(response.Data.([]string), email)
			continue
		}

		err = groupsrepo.SendInvitation(idUser, inv.IdGroup)

		if err != nil {
			return nil, services.ErrorInternal
		}
	}

	if len(response.Data.([]string)) > 0 {
		return &response, nil
	}

	return nil, nil
}
