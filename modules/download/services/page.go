/*
 * @Author: liziwei01
 * @Date: 2023-05-11 19:37:28
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-05-11 23:12:43
 * @Description: file content
 */
package services

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gogf/gf/util/gconv"
	"github.com/liziwei01/gin-lib/library/env"
	"github.com/liziwei01/gin-monitor-appui/modules/download/constants"
	downloadModel "github.com/liziwei01/gin-monitor-appui/modules/download/model"
)

func Page(ctx context.Context, pars downloadModel.PagePars) (map[string]interface{}, error) {
	files, err := filesPath(pars.Path)
	if err != nil {
		return nil, err
	}

	filesLen := gconv.Uint(len(files))
	start := pars.Page * pars.PageLength
	end := (pars.Page + 1) * pars.PageLength
	if start > filesLen {
		return nil, nil
	}
	if end > filesLen {
		end = filesLen
	}

	return map[string]interface{}{
		"list":  filesBase(files[start:end]),
		"count": filesLen,
	}, nil
}

func filesPath(path string) ([]string, error) {
	var files []string
	absPath := filepath.Join(env.Default.RootDir(), constants.DATA_LOCATION, path)
	err := filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return err
	})
	return files, err
}

func filesBase(files []string) []string {
	for i := 0; i != len(files); i++ {
		files[i] = filepath.Base(files[i])
	}

	return files
}
