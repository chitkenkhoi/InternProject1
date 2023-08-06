package User

import (
	"github.com/gin-gonic/gin"
	"project/API/Email"
	"project/DB"
	"project/model"
	"strconv"
	s "strings"
)

func CountUserInternal(ctx *gin.Context) {
	var count int64
	DB.DB.Where("status = ?", "Internal").Table("users").Count(&count)
	ctx.JSON(200, gin.H{
		"count": count,
	})
}
func CountUserPending(ctx *gin.Context) {
	var count int64
	DB.DB.Where("status = ?", "Pending").Table("users").Count(&count)
	ctx.JSON(200, gin.H{
		"count": count,
	})
}
func GetUser(ctx *gin.Context) {

	var user []model.User
	if err := DB.DB.Order("id").Table("users").Find(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": user,
	})

}
func GenerateID(ctx *gin.Context) {
	var user model.User
	if err := DB.DB.Table("users").Last(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"ID": user.Id,
	})

}

func GetUserInternal(ctx *gin.Context) {

	var page, _ = strconv.Atoi(ctx.Param("page"))
	if page == 0 {
		var count int64
		DB.DB.Where("status = ?", "Internal").Order("id").Table("users").Count(&count)
		page = int((count-count%8)/8 + 1)
		if count%8 == 0 && count != 0 {
			page--
		}

	}
	page = (page - 1) * 8
	var user []model.User
	if err := DB.DB.Where("status = ?", "Internal").Order("id").Table("users").Limit(8).Offset(page).Find(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": user,
		"page": page/8 + 1,
	})
}
func GetUserPending(ctx *gin.Context) {
	var page, _ = strconv.Atoi(ctx.Param("page"))
	if page == 0 {
		var count int64
		DB.DB.Where("status = ?", "Pending").Order("id").Table("users").Count(&count)
		page = int((count-count%8)/8 + 1)
		if count%8 == 0 && count != 0 {
			page--
		}

	}
	page = (page - 1) * 8
	var user []model.User
	if err := DB.DB.Where("status = ?", "Pending").Order("id").Table("users").Limit(8).Offset(page).Find(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": user,
		"page": page/8 + 1,
	})
}
func SearchUserPending(ctx *gin.Context) {
	var data = ctx.Param("param")
	var page, _ = strconv.Atoi(ctx.Param("page"))
	if page == 0 {
		var count int64
		DB.DB.Where("status = ? AND ((first_name || ' ' || last_name) LIKE ? OR email LIKE ?)", "Pending", "%"+data+"%", "%"+data+"%").Order("id").Table("users").Count(&count)
		page = int((count-count%8)/8 + 1)
		if count%8 == 0 && count != 0 {
			page--
		}
	}
	page = (page - 1) * 8
	var users []model.User
	if err := DB.DB.Where("status = ? AND ((first_name || ' ' || last_name) LIKE ? OR email LIKE ?)", "Pending", "%"+data+"%", "%"+data+"%").Order("id").Table("users").Limit(8).Offset(page).Find(&users).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{
		"data": users,
		"page": page/8 + 1,
	})

}
func SearchUserInternal(ctx *gin.Context) {
	var data = ctx.Param("param")
	var page, _ = strconv.Atoi(ctx.Param("page"))
	if page == 0 {
		var count int64
		DB.DB.Where("status = ? AND ((first_name || ' ' || last_name) LIKE ? OR email LIKE ?)", "Internal", "%"+data+"%", "%"+data+"%").Order("id").Table("users").Count(&count)
		page = int((count-count%8)/8 + 1)
		if count%8 == 0 && count != 0 {
			page--
		}
	}
	page = (page - 1) * 8
	var users []model.User
	if err := DB.DB.Where("status = ? AND ((first_name || ' ' || last_name) LIKE ? OR email LIKE ?)", "Internal", "%"+data+"%", "%"+data+"%").Order("id").Table("users").Limit(8).Offset(page).Find(&users).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{
		"data": users,
		"page": page/8 + 1,
	})

}
func GetUserByid(ctx *gin.Context) {
	id := ctx.Param("id")
	var user model.User
	if err := DB.DB.Where("id = ?", id).Table("users").First(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": user,
	})

}
func GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	var user model.User
	if err := DB.DB.Where(" status = ? AND email = ?", "Internal", email).Table("users").First(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": user,
	})

}
func ResendEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	lst := []string{}
	lst = append(lst, email)
	subject := "Resend invitation"
	body := "Click the link below to sign up and become a member of us:\nLink: http://localhost:8081/auth/signup"
	if Email.SendEmail(lst, subject, body) {
		ctx.JSON(200, gin.H{
			"isSuccess": true,
		})
	} else {
		ctx.JSON(200, gin.H{
			"isSuccess": false,
		})
	}

}

func CreateUser(ctx *gin.Context) {
	var users []model.User
	ctx.ShouldBind(&users)

	if err := DB.DB.Table("users").Create(&users).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	var list_recipient = []string{}
	var size = len(users)
	for i := 0; i < size; i++ {
		list_recipient = append(list_recipient, users[i].Email)
	}
	body := "Click the link below to sign up and become a member of us:\nLink: http://localhost:8081/auth/signup "
	subject := "Invite new member"

	if Email.SendEmail(list_recipient, subject, body) {
		ctx.JSON(200, gin.H{
			"data": users,
		})
	} else {
		ctx.JSON(200, gin.H{
			"err": "got an error",
		})
	}
}
func UpdateUserByid(ctx *gin.Context) {
	id := ctx.Param("id")
	var user model.User
	ctx.ShouldBind(&user)
	if err := DB.DB.Where("id = ?", id).Table("users").Save(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": user,
	})
}
func UploadAvatarUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	dst := "C:/Project/my-project/public/img/avatar" + "_" + id + ".jpg"
	var file model.File
	file.File, _ = ctx.FormFile("file")
	ctx.SaveUploadedFile(file.File, dst)

	if err := DB.DB.Model(&model.User{}).Where("id = ?", id).Update("avatar", s.Replace(dst, "C:/Project/my-project/public", "", 1)).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": s.Replace(dst, "C:/Project/my-project/public", "", 1),
	})
}
func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user model.User
	if err := DB.DB.Where("id = ?", id).Table("users").Delete(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": user,
	})
}
