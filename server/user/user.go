package user

type User struct {
	Id       string
	Username string `json:"username"`
	Password string `json:"password"`
}
