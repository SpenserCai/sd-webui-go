/*
 * @Author: SpenserCai
 * @Date: 2023-08-13 12:33:39
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-13 21:53:49
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type RembgRequest = SdApiModel.BodyRembgRemoveRembgPost

// Checked: True
type RembgResponse struct {
	Image string `json:"image"`
}
