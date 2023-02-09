package util

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/go-gin-example/pkg/setting"
	"github.com/unknwon/com"
)

//分页逻辑

func GetPage(c *gin.Context) int {
	ret := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		//分页信息
		ret = (page - 1) * setting.PageSize
	}
	return ret
}
