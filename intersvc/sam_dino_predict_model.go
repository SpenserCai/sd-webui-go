/*
 * @Author: SpenserCai
 * @Date: 2023-08-14 01:00:25
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-29 20:38:25
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type SamDinoPredictRequest = SdApiModel.DINOPredictRequest

// Checked: True
type SamDinoPredictResponse struct {
	Msg          string   `json:"msg"`
	ImageWithBox []string `json:"image_with_box"`
}
