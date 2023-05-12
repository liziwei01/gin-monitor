/*
 * @Author: liziwei01
 * @Date: 2023-05-11 19:26:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-05-11 20:09:38
 * @Description: file content
 */
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liziwei01/gin-lib/library/response"
	downloadModel "github.com/liziwei01/gin-monitor-appui/modules/download/model"
	downloadService "github.com/liziwei01/gin-monitor-appui/modules/download/services"
)

func Page(ctx *gin.Context) {
	inputs, hasError := getPagePars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}

	res, err := downloadService.Page(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}

	response.StdSuccess(ctx, res)
}

func getPagePars(ctx *gin.Context) (downloadModel.PagePars, bool) {
	var inputs downloadModel.PagePars

	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}

	if inputs.PageLength == 0 {
		inputs.PageLength = 10
	}

	return inputs, false
}
