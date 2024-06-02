package learn_golang_gorm

import "time"

type User struct {
	ID       int
	Name     Name `gorm:"embedded"`
	Email    string
	Password string
	DateAt   DateAt `gorm:"embedded"`
}

type Name struct {
	FirstName string
	LastName  string
}

type DateAt struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
