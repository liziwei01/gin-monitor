/*
 * @Author: liziwei01
 * @Date: 2023-05-09 22:58:00
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-05-11 02:53:24
 * @Description: file content
 */

package httpapi

import (
	"net/http"

	"github.com/liziwei01/gin-lib/middleware"
	downloadRouters "github.com/liziwei01/gin-monitor-appui/modules/download/routers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: start http server and start listening
 * @param {*}
 * @return {*}
 */
func InitRouters(handler *gin.Engine) {
	//暂时解决跨域问题
	handler.Use(middleware.CrossRegionMiddleware())

	// init routers
	router := handler.Group("/api")
	downloadRouters.Init(router)

	// safe router
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello! THis is iMonitor.")
	})
}
