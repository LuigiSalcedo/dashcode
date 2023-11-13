package models

// Login Model
type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User Model
type UserModel struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// Full User Model
type FullUserModel struct {
	UserModel
	LoginModel
}

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"password"`
}
