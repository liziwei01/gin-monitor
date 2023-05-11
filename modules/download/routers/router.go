/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-05-11 01:00:27
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
	}
	// pageGroup := router.Group("/page")
	// {
	// 	pageGroup.GET("/", downloadController.StreamLocal)
	// }
}
