package routes

import (
	"github.com/ZhangYongChang/learngo/ginSample/controller"
	"github.com/ZhangYongChang/learngo/ginSample/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleware(), controller.QueryUser)
	return r
}
