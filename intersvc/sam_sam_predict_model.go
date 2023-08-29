/*
 * @Author: SpenserCai
 * @Date: 2023-08-14 01:00:25
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-29 20:38:33
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type SamSamPredictRequest = SdApiModel.SamPredictRequest

// Checked: True
type SamSamPredictResponse struct {
	Msg           string   `json:"msg"`
	BlendedImages []string `json:"blended_images"`
	Masks         []string `json:"masks"`
	MaskedImages  []string `json:"masked_images"`
}
