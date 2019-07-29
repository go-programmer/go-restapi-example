package user

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type Users []User
