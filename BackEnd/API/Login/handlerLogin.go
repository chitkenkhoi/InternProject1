package Login

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"project/API/Email"
	"project/DB"
	"project/model"
)

func CheckEmail(email string) (string, string) {
	var user model.User

	if err := DB.DB.Where("email = ? AND status = ?", email, "Internal").Table("users").First(&user).Error; err != nil {
		if er := DB.DB.Where("email = ? AND status = ?", email, "Pending").Table("users").First(&user).Error; er != nil {
			return er.Error(), "error"
		}
		return user.Id, "foundPending"
	} else {
		return user.Id, "foundInternal"
	}

}
func Login(ctx *gin.Context) {
	var loginRequest model.LoginRequest
	ctx.ShouldBind(&loginRequest)
	id, code := CheckEmail(loginRequest.UserLogin)
	if code == "foundPending" || code == "error" {
		ctx.JSON(200, gin.H{
			"err": code,
		})
	} else {
		var login model.Login
		if err := DB.DB.Where("user_id = ? AND user_password = ?", id, loginRequest.Password).Table("login").First(&login).Error; err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
			return
		}
		const str = "0123456789"
		var OTP = ""
		for i := 0; i < 6; i++ {
			OTP = OTP + string(str[rand.Intn(10)])
		}
		var twoFA model.TwoFA
		DB.DB.Table("2FA").Where("email = ?", loginRequest.UserLogin).Delete(&twoFA)
		twoFA.Email = loginRequest.UserLogin
		twoFA.OtpCode = OTP

		DB.DB.Table("2FA").Create(&twoFA)
		const subject = "Confirm Login"
		var body = "Your verification code is: " + OTP + "\nPlease complete the account verification process."
		ctx.JSON(200, gin.H{
			"data": Email.SendEmail([]string{loginRequest.UserLogin}, subject, body),
		})
	}

}
func ResetPasswordEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	var user model.User
	if err := DB.DB.Where("email = ? AND status = ?", email, "Internal").Table("users").First(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	var listEmail = []string{}
	listEmail = append(listEmail, email)
	subject := "Reset password confirming"
	body := "You have made a request to reset your password. If you did not do it, just ignore this email. Otherwise, please click the link below:\nLink: http://localhost:8081/auth/emailConfirm/" + email
	ctx.JSON(200, gin.H{
		"isSuccess": Email.SendEmail(listEmail, subject, body),
	})
}
func CreateLogin(ctx *gin.Context) {
	var loginRequest model.LoginRequest
	ctx.ShouldBind(&loginRequest)
	id, code := CheckEmail(loginRequest.UserLogin)
	if code == "foundPending" {
		if err := DB.DB.Table("users").Where("id = ?", id).Update("status", "Internal").Error; err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
			return
		}

		var login model.Login
		login.User_password = loginRequest.Password
		login.User_id = id
		if err := DB.DB.Table("login").Create(&login).Error; err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
			return
		}
		const subject = "Confirm Signup"
		const body = "You are signing up for our website, please verify your action. If you did not do this, please contact security team or reset your account's password. "
		if Email.SendEmail([]string{loginRequest.UserLogin}, subject, body) {
			ctx.JSON(200, gin.H{
				"data": login,
			})
		}

	} else {
		ctx.JSON(200, gin.H{
			"err": code,
		})
	}

}
func UpdatePassword(ctx *gin.Context) {
	var login_request model.LoginRequest
	ctx.ShouldBind(&login_request)
	id, code := CheckEmail(login_request.UserLogin)
	if code == "foundInternal" {
		var login model.Login
		login.User_password = login_request.Password
		login.User_id = id
		if err := DB.DB.Where("user_id = ?", id).Table("login").Save(&login).Error; err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": login,
		})
	} else {
		ctx.JSON(200, gin.H{
			"err": code,
		})
	}

}
func Handle2FA(ctx *gin.Context) {
	var twoFA model.TwoFA
	ctx.ShouldBind(&twoFA)
	otp := twoFA.OtpCode
	if err := DB.DB.Where("email = ?", twoFA.Email).Table("2FA").First(&twoFA).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	if otp != twoFA.OtpCode {
		ctx.JSON(200, gin.H{
			"err": "Your code is incorrect.",
		})
		return
	}
	DB.DB.Table("2FA").Where("email = ?", twoFA.Email).Delete(&twoFA)
	ctx.JSON(200, gin.H{
		"data": true,
	})
}
