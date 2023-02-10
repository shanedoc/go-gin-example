package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/go-gin-example/middleware/jwt"
	"github.com/shanedoc/go-gin-example/pkg/setting"
	"github.com/shanedoc/go-gin-example/routes/api"
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

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	//jwt中间件
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTage)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)          //列表
		apiv1.GET("/articles/:id", v1.GetArticle)       //详情
		apiv1.POST("/article", v1.AddArticle)           //新增
		apiv1.PUT("/article/:id", v1.EditArticle)       //编辑
		apiv1.DELETE("/articles/:id", v1.DeleteArticle) //删除

	}

	return r
}
