package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

func SetupRouter(DB *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.GET("getJob", func(c *gin.Context) {
		InfoTransfer(c, DB)
		//c.String(http.StatusOK, "获取所有工作")
		log.Printf("获取所有工作")
	})

	return r
}
