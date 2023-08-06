package model

type Login struct {
	User_id       string `json:"user_id" gorm:"column:user_id"`
	User_password string `json:"user_password" gorm:"column:user_password"`
}
type LoginRequest struct {
	UserLogin string `json:"userEmail"`
	Password  string `json:"password"`
}
type TwoFA struct {
	Email   string `json:"email" gorm:"column:email"`
	OtpCode string `json:"otpCode" gorm:"column:otpCode"`
}
