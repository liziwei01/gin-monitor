/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:15:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-05-11 19:30:06
 * @Description: file content
 */
package model

type StreamPars struct {
	Path string `json:"path" form:"path" binding:"required"`
}
