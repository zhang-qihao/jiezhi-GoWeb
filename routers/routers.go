package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

func SetupRouter(DB *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.GET("getJob", func(c *gin.Context) {
		rank := GetJobInfoFromDatabase(DB, "")
		InfoTransfer(c, rank)
		//c.String(http.StatusOK, "获取所有工作")
		log.Printf("获取所有工作")
	})

	r.GET("queryJob", func(c *gin.Context) {
		query := c.DefaultQuery("query", "")
		// 获取要放在主页的职位信息
		rank := GetJobInfoFromDatabase(DB, query)
		// 向前端传JSON数据
		InfoTransfer(c, rank)
	})

	r.GET("getJobDetail", func(c *gin.Context) {
		query := c.DefaultQuery("jid", "")
		// 获取要放在主页的职位信息
		rank := GetJobDetailFromDatabase(DB, query)
		// 向前端传JSON数据
		InfoTransfer(c, rank)
	})

	return r
}
