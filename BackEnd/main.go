package main

import (
	"github.com/gin-contrib/cors"
	"project/API/Company"
	"project/API/Login"
	"project/API/User"

	"project/DB"

	"github.com/gin-gonic/gin"
)

func main() {
	DB.CreateDB()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:8081"},
		AllowMethods:  []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:  []string{"Content-Type", "Content-Length", "Accept-Encoding"},
		ExposeHeaders: []string{"Content-Length", "Content-Type"},
	}))
	api := router.Group("/api")
	{
		api.GET("getUser/", User.GetUser)
		api.GET("countUser/Internal", User.CountUserInternal)
		api.GET("countUser/Pending", User.CountUserPending)
		api.GET("getUser/Internal/:page/", User.GetUserInternal)
		api.GET("getUser/Pending/:page/", User.GetUserPending)
		api.GET("getUser/Pending/:page/:param", User.SearchUserPending)
		api.GET("generateID/", User.GenerateID)
		api.GET("getUser/Internal/:page/:param", User.SearchUserInternal)
		api.GET("getUser/:email", User.GetUserByEmail)
		api.GET("getUserByid/:id", User.GetUserByid)
		api.GET("resendEmail/:email/", User.ResendEmail)
		api.GET("getCompanyInformation/:control_number", Company.GetCompanyByControlNumber)
		api.GET("resetPassword/:email", Login.ResetPasswordEmail)
		api.POST("createUser", User.CreateUser)
		api.POST("createCompany", Company.CreateCompany)
		api.POST("createLogin", Login.CreateLogin)
		api.POST("Login", Login.Login)
		api.POST("checkTwoFA", Login.Handle2FA)
		api.PUT("/uploadLogo", Company.UploadLogo)
		api.PUT("/uploadSignature", Company.UploadSignature)
		api.PUT("updateCompany/:control_number", Company.UpdateCompanyByControlNumber)
		api.PUT("updateUser/:id", User.UpdateUserByid)
		api.PUT("uploadUserAvatar/:id", User.UploadAvatarUserById)
		api.PUT("updatePassword/", Login.UpdatePassword)
		api.DELETE("deleteUser/:id", User.DeleteUser)

	}
	router.Run(":8080")
}
