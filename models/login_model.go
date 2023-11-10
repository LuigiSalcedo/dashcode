package models

// Login Structure
type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// Full User Model
type FullUserModel struct {
	UserModel
	LoginModel
}
