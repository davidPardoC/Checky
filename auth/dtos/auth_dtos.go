package dtos

type LoginDto struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type SignupDto struct {
	Name     string `json:"name"  binding:"required"`
	LastName string `json:"lastName"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
