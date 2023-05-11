/*
 * @Author: liziwei01
 * @Date: 2023-05-09 22:57:00
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-05-09 22:57:00
 * @Description: file content
 */
package main

import (
	"log"

	"github.com/liziwei01/gin-lib/bootstrap"
	"github.com/liziwei01/gin-monitor-appui/httpapi"
)

func main() {
	app, err := bootstrap.Setup()
	if err != nil {
		log.Fatalln(err)
	}
	// 注册接口路由
	httpapi.InitRouters(app.Handler)

	app.Start()
}
