/*
 * @Author: SpenserCai
 * @Date: 2023-08-14 01:00:25
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-14 13:10:11
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type SamDinoPredictRequest = SdApiModel.DINOPredictRequest

type SamDinoPredictResponse struct {
	Msg          string   `json:"msg"`
	ImageWithBox []string `json:"image_with_box"`
}
