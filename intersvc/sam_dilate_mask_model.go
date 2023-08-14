/*
 * @Author: SpenserCai
 * @Date: 2023-08-14 01:00:25
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-14 13:18:18
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type SamDilateMaskRequest = SdApiModel.DilateMaskRequest

type SamDilateMaskResponse struct {
	BlendedImage string `json:"blended_image"`
	Mask         string `json:"mask"`
	MaskedImage  string `json:"masked_image"`
}
