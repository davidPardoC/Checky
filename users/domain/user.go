package domain

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
}
