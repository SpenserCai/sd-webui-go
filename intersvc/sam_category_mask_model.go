/*
 * @Author: SpenserCai
 * @Date: 2023-08-14 01:00:25
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-14 13:30:37
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type SamCategoryMaskRequest = SdApiModel.BodyAPICategoryMaskSamCategoryMaskPost

type SamCategoryMaskResponse struct {
	Msg          string `json:"msg"`
	BlendedImage string `json:"blended_image"`
	Mask         string `json:"mask"`
	MaskedImage  string `json:"masked_image"`
	ResizeInput  string `json:"resize_input"`
}
