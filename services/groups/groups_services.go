package groups

import (
	"dashcode/repositories/groupsrepo"
	"dashcode/services"
)

// Services to create a group
func CreateGroup(idCreator int64, name, description string) *services.Error {
	err := groupsrepo.CreateGroup(idCreator, name, description)

	if err != nil {
		return services.ErrorBadRequest
	}

	return nil
}
