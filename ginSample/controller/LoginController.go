package controller

import (
	"net/http"

	"github.com/ZhangYongChang/learngo/ginSample/common"
	"github.com/ZhangYongChang/learngo/ginSample/model"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	db := common.GetDB()
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

	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	if user.Password != password {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户密码错误"})
		return
	}

	ctx.JSON(200, gin.H{"msg": "登录成功"})
}
