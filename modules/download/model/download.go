/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:15:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-05-11 02:52:00
 * @Description: file content
 */
package model

type StreamPars struct {
	Path string `form:"path" json:"path" binding:"required"`
}
