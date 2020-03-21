package controller

import (
	"net/http"

	"github.com/ZhangYongChang/learngo/ginSample/dto"
	"github.com/ZhangYongChang/learngo/ginSample/model"
	"github.com/gin-gonic/gin"
)

func QueryUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}
