/*
 * @Author: liziwei01
 * @Date: 2023-10-05 00:01:52
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-05 01:29:05
 * @Description: file content
 */
package controllers

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liziwei01/gin-lib/library/env"
	"github.com/liziwei01/gin-lib/library/response"
	"github.com/liziwei01/gin-lib/library/utils"
	"github.com/liziwei01/gin-monitor-appui/modules/download/constants"
)

const (
	MD5DIR = "md5"
	// 资源类型
	HEADER_TYPE_INFO = "Content-Type"
)

func Upload(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		response.StdInvalidParams(ctx)
		return
	}

	// 获取文件字节流
	byteStream, err := utils.File.GetFileBytes(fileHeader)
	if err != nil {
		err = fmt.Errorf("获取文件字节流失败: %s", err.Error())
		response.StdFailed(ctx, err)
	}

	filePath := filepath.Join(env.Default.RootDir(), constants.DATA_LOCATION)
	fileName := genFileName(filePath, fileHeader)
	os.WriteFile(fileName, byteStream, 0644)

	response.StdSuccess(ctx, fileName)
}

// 生成完整文件名（包含绝对路径）
func genFileName(filePath string, fileHeader *multipart.FileHeader) string {
	// 获取文件后缀
	fileSubfix := genFileSubfix(fileHeader.Header.Get(HEADER_TYPE_INFO))

	// 生成唯一文件名
	fileName := utils.UUID.GenUUID() + "." + fileSubfix

	// 返回完整文件名
	return path.Join(filePath, fileName)
}

// 根据 fileType 获取文件后缀  如 'video/mp4'-> 'mp4'
func genFileSubfix(fileType string) string {
	subfixIdx := strings.LastIndex(fileType, "/") + 1
	return fileType[subfixIdx:]
}
