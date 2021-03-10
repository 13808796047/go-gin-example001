package handlers

import (
	"net/http"

	"github.com/13808796047/go-gin-example/models"
	"github.com/13808796047/go-gin-example/pkg/e"
	"github.com/13808796047/go-gin-example/pkg/setting"
	"github.com/13808796047/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取多个文章标签
func Index(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
		maps["name"] = state
	}

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.GetMsg(e.SUCCESS),
		"data":    data,
	})
}

//新增文章标签
func Store(c *gin.Context) {
}

//修改文章标签
func Update(c *gin.Context) {
}

//删除文章标签
func Destory(c *gin.Context) {
}
