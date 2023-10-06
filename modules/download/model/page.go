/*
 * @Author: liziwei01
 * @Date: 2022-07-02 19:42:48
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-05 00:54:48
 * @Description: file content
 */
package model

type PagePars struct {
	Page       uint `json:"page" form:"page"`
	PageLength uint `json:"page_length" form:"page_length"`

	Path string `json:"path" form:"path"`
}
