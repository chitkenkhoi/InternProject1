package Company

import (
	"github.com/gin-gonic/gin"
	"project/DB"
	"project/model"
)

func GetCompanyByControlNumber(ctx *gin.Context) {
	control_number := ctx.Param("control_number")
	var company model.Company
	if err := DB.DB.Where("control_number = ?", control_number).Table("company").First(&company).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": company,
	})

}
func CreateCompany(ctx *gin.Context) {
	var company model.Company
	ctx.ShouldBind(&company)

	if err := DB.DB.Table("company").Create(&company).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": company,
	})
}
func UploadLogo(ctx *gin.Context) {
	var file model.File
	file.File, _ = ctx.FormFile("file")
	dst := "C:/Project/my-project/public/img/logo.jpg"
	ctx.SaveUploadedFile(file.File, dst)

}
func UploadSignature(ctx *gin.Context) {
	var file model.File
	file.File, _ = ctx.FormFile("file")
	dst := "C:/Project/my-project/public/img/sign.jpg"
	ctx.SaveUploadedFile(file.File, dst)

}
func UpdateCompanyByControlNumber(ctx *gin.Context) {
	control_number := ctx.Param("control_number")
	var company model.CompanyBody
	ctx.ShouldBind(&company)
	if err := DB.DB.Where("control_number = ?", control_number).Table("company").Save(&company.Body).Error; err != nil {
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": company.Body,
	})

}
