package controller

import (
	"log"
	"net/http"

	"github.com/ZhangYongChang/learngo/ginSample/common"
	"github.com/ZhangYongChang/learngo/ginSample/model"
	"github.com/ZhangYongChang/learngo/ginSample/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	db := common.GetDB()
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		util.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		util.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}

	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		util.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		util.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户密码错误")
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		util.Response(ctx, http.StatusInternalServerError, 500, nil, "系统错误")
		log.Printf("token generate error : %v", err)
		return
	}
	util.Success(ctx, gin.H{"token": token}, "登录成功")
}
