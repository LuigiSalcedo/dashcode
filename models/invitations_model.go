package models

type InvitationData struct {
	Id               int64  `json:"id"`
	GroupId          int64  `json:"id_group"`
	GroupName        string `json:"name_group"`
	GroupDescription string `json:"description_group"`
	CreatorName      string `json:"creator_name"`
}
