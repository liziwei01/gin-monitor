/*
 * @Author: liziwei01
 * @Date: 2023-05-11 02:55:01
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-05-11 20:06:14
 * @Description: file content
 */
package middleware

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/liziwei01/gin-lib/library/env"
	"github.com/liziwei01/gin-lib/library/utils"
	"github.com/liziwei01/gin-monitor-appui/library/totp"
)

const path = "secret"

var secrets []string

func init() {
	var files []string
	absPath := filepath.Join(env.Default.RootDir(), path)
	filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	for i := 0; i != len(files); i++ {
		sc := readFile(files[i])
		if sc != "" {
			secrets = append(secrets, sc)
		}
	}
}

func CheckCodeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		code, ok := utils.Request.Header(c.Request, "code")
		if isRealease() != true {
			// 线下无限制.
			c.Next()
		} else if !ok {
			c.Abort()
		} else if !checkCodes(code) {
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func checkCodes(code string) bool {
	for i := 0; i != len(secrets); i++ {
		if totp.Validate(secrets[i], code) {
			return true
		}
	}
	return false
}

func readFile(absPath string) string {
	file, err := os.Open(absPath)
	if err != nil {
		return ""
	}
	byteStream, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(byteStream)
}

// 判断是否为线上环境.
func isRealease() bool {
	return env.RunMode() == env.RunModeRelease
}