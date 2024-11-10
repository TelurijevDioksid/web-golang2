package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SqlInjUser struct {
	Inj      string `json:"inj"`
	Username string `json:"username"`
	Password string `json:"password"`
}
