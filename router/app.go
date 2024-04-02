package router

import "github.com/gin-gonic/gin"
import "cloud-disk/service"

func Router() *gin.Engine {

	r := gin.Default()

	//用户模块
	r.POST("/user/login", service.Login)

	return r
}
