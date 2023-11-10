package models

// Create group
type CreateGroup struct {
	IdCreator   int64  `json:"id_creator"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Full Group
type Group struct {
	Id int64 `json:"id"`
	CreateGroup
}
