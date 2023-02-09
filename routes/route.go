package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/go-gin-example/pkg/setting"
	v1 "github.com/shanedoc/go-gin-example/routes/api/v1"
)

//路由文件

func InitRouter() *gin.Engine {
	//初始化gin实例
	r := gin.Default()
	//使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTage)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

	}

	return r
}
