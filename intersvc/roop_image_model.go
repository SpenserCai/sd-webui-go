/*
 * @Author: SpenserCai
 * @Date: 2023-08-22 11:34:42
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-29 20:38:06
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type RoopImageRequest = SdApiModel.BodyRoopImageRoopImagePost

// Checked: True
type RoopImageResponse struct {
	Image string `json:"image"`
}
