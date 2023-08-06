package model

import (
	"time"
)

type Company struct {
	Control_number      string     `json:"control_number" gorm:"column:control_number"`
	C_name              *string    `json:"c_name" gorm:"c_name"`
	Logo                *string    `json:"logo" gorm:"logo"`
	Address             *string    `json:"address" gorm:"column:address"`
	Phone_number        *string    `json:"phone_number" gorm:"column:phone_number"`
	Representative_name *string    `json:"representative_name" gorm:"column:representative_name"`
	Partner_company     *string    `json:"partner_company" gorm:"column:partner_company"`
	Signature           *string    `json:"signature" gorm:"column:signature"`
	TwoFA               bool       `json:"twoFA" gorm:"column:twoFA"`
	Email_notification  bool       `json:"email_notification" gorm:"column:email_notification"`
	Email_destination   string     `json:"email_destination" gorm:"column:email_destination"`
	Created_at          *time.Time `json:"created_at" gorm:"column:created_at"`
}

type CompanyBody struct {
	Body Company `json:"body"`
}
