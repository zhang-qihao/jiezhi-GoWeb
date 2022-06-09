package routers

import (
	"MySystem/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	v1Group := r.Group("v1")
	{
		// 获取所有工作
		v1Group.GET("getJob", controller.GetJob)
		// 查找工作
		v1Group.GET("queryJob", controller.GetJob)
		// 查询工作详情
		v1Group.GET("getJobDetail", controller.GetJobDetail)
		// 新增工作
		v1Group.POST("publishJob", controller.PublishAJob)
	}

	return r
}
