package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Joined   string `json:"joined"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserWoPass struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Joined   string `json:"joined"`
	IsAdmin  bool   `json:"is_admin"`
}

type NewUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Joined   string `json:"joined"`
	IsAdmin  bool   `json:"is_admin"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
