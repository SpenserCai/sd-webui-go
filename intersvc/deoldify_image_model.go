/*
 * @Author: SpenserCai
 * @Date: 2023-08-12 15:21:45
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-13 21:42:04
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type DeoldifyImageRequest = SdApiModel.BodyDeoldifyImageDeoldifyImagePost

// Checked: True
type DeoldifyImageResponse struct {
	Image string `json:"image"`
}
