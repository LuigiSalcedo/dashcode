package groups

import (
	"dashcode/models"
	"dashcode/repositories/groupsrepo"
	"dashcode/services"
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
