package route

import (
	"example/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("login", service.Login)
		authGroup.POST("sendCode", service.SendVerifyCode)
		authGroup.POST("register", service.Register)
	}
	userGroup := r.Group("/api/user", preHandler())
	{
		userGroup.GET("myInfo", service.GetMyInfo)
	}

	profileGroup := r.Group("/api/profile", preHandler())
	{
		profileGroup.POST("create", service.CreateProfile)
		profileGroup.GET("get", service.GetProfile)
		profileGroup.GET("/:id/get", service.GetProfileByID)
		profileGroup.PUT("update", service.UpdateProfile)
		profileGroup.DELETE("delete", service.DeleteProfile)
	}
}
