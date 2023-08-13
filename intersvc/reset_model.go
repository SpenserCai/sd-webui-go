/*
 * @Author: SpenserCai
 * @Date: 2023-08-13 12:33:39
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-13 21:54:01
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type ResetRequest = SdApiModel.ResetBody

// Checked: True
type ResetResponse struct {
	Success bool `json:"success"`
}
