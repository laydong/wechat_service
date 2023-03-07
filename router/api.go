package router

import (
	"github.com/gin-gonic/gin"
	"wechat-service/handler"
)

func ApiRouter(r *gin.Engine) {

	v := r.Group("/api/")
	{

		v.POST("test", handler.Test)
	}
}
