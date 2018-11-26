package entity

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) IsAuthenticated(password string) bool {
	return u.Password == password
}
