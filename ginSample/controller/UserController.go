package controller

import (
	"net/http"

	"github.com/ZhangYongChang/learngo/ginSample/common"
	"github.com/ZhangYongChang/learngo/ginSample/model"
	"github.com/ZhangYongChang/learngo/ginSample/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于6位"})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	if isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)

	ctx.JSON(200, gin.H{"msg": "注册成功"})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
