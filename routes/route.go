package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/go-gin-example/pkg/setting"
)

//路由文件

func InitRouter() *gin.Engine {
	//初始化gin实例
	r := gin.Default()
	//使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
	return r
}
