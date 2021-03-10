package handlers

import (
	"github.com/gin-gonic/gin"
)

//获取多个文章标签
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Index",
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
