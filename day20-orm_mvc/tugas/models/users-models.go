package models

import "time"

type UsersGorm struct {
	ID    int       `gorm:"column:id"`
	Email string    `gorm:"column:email"`
	Name  string    `gorm:"column:name"`
	DOB   time.Time `gorm:"column:dob"`
}

func (UsersGorm) TableName() string {
	return "users"
}

type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	DOB   string `json:"dob,omitempty"`
}
