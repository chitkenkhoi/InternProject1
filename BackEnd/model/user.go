package model

import (
	"time"
)

type User struct {
	Id           string     `json:"id" gorm:"column:id;"`
	Email        string     `json:"email" gorm:"column:email"`
	Role         string     `json:"role" gorm:"column:role"`
	Avatar       *string    `json:"avatar" gorm:"column:avatar"`
	First_name   *string    `json:"first_name" gorm:"column:first_name"`
	Last_name    *string    `json:"last_name" gorm:"column:last_name"`
	Director     *string    `json:"director" gorm:"column:director"`
	Phone_number *string    `json:"phone_number" gorm:"column:phone_number"`
	Created_at   *time.Time `json:"created_at" gorm:"column:created_at"`
	Status       string     `json:"status" gorm:"column:status"`
	Work_for     string     `json:"work_for" gorm:"column:work_for"`
}
