/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-05 00:36:31
 * @Description: file content
 */
package routers

import (
	"github.com/gin-gonic/gin"

	downloadController "github.com/liziwei01/gin-monitor-appui/modules/download/controllers"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.RouterGroup) {
	downloadGroup := router.Group("/download")
	// apiGroup.Use(middleware.CheckLoginMiddleware())
	{
		downloadGroup.GET("/stream", downloadController.Stream)
		downloadGroup.GET("/page", downloadController.Page)
		downloadGroup.POST("/upload", downloadController.Upload)
	}
}
