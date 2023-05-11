/*
 * @Author: liziwei01
 * @Date: 2023-05-09 23:00:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-05-11 02:53:31
 * @Description: file content
 */
package controllers

import (
	"github.com/liziwei01/gin-lib/library/response"
	downloadModel "github.com/liziwei01/gin-monitor-appui/modules/download/model"

	"github.com/gin-gonic/gin"
)

const (
	location = "data"
)

func Stream(ctx *gin.Context) {
	inputs, hasError := getStreamPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}

	if len(inputs.Path) == 0 {
		inputs.Path = "/"
	} else if inputs.Path[0] != '/' {
		inputs.Path = "/" + inputs.Path
	}

	absPath := location + inputs.Path
	ctx.File(absPath)
}

func getStreamPars(ctx *gin.Context) (downloadModel.StreamPars, bool) {
	var inputs downloadModel.StreamPars

	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}

	return inputs, false
}
