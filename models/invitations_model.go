package models

type InvitationData struct {
	Id               int64  `json:"id"`
	GroupId          int64  `json:"id_group"`
	GroupName        string `json:"name_group"`
	GroupDescription string `json:"description_group"`
	CreatorName      string `json:"creator_name"`
}

type SentInvitationsData struct {
	Id        int64  `json:"id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	State     string `json:"state"`
}

type UserInvitationData struct {
	Id               int64  `json:"id"`
	GroupId          int64  `json:"group_id"`
	GroupName        string `json:"group_name"`
	GroupDescription string `json:"group_description"`
}
