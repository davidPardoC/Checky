package domain

type Role struct {
	Name string `json:"name" gorm:"unique;not null"`
}
