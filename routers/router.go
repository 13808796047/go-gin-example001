package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/13808796047/go-gin-example/handlers"
	"github.com/13808796047/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	api := r.Group("/api/v1")
	{
		// 获取标签列表
		api.GET("/tags", handlers.Index)
	}

	return r
}
