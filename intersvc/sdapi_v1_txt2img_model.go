/*
 * @Author: SpenserCai
 * @Date: 2023-08-13 12:33:39
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-13 13:17:58
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type SdapiV1Txt2imgRequest = SdApiModel.StableDiffusionProcessingTxt2Img

type SdapiV1Txt2imgResponse struct {
	Images []string `json:"images"`
	Info   string   `json:"info"`
}
