/*
 * @Author: SpenserCai
 * @Date: 2023-08-13 12:33:39
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-13 13:10:22
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type RembgRequest = SdApiModel.BodyRembgRemoveRembgPost

type RembgResponse struct {
	Image string `json:"image"`
}
