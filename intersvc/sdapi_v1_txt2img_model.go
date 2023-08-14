/*
 * @Author: SpenserCai
 * @Date: 2023-08-13 12:33:39
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-14 13:01:44
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type SdapiV1Txt2imgRequest = SdApiModel.StableDiffusionProcessingTxt2Img

// Checked: True
type SdapiV1Txt2imgResponse = SdApiModel.TextToImageResponse
