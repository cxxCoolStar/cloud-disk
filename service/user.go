package service

import (
	"cloud-disk/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginRequest := models.LoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{
			"code":    -1,
			"message": "参数错误",
		})
		return
	}

	fmt.Print(loginRequest.Username, "  >>>>>>>>>>>  ", loginRequest.Password)
}
