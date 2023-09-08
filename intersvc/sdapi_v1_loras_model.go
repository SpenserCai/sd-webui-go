/*
 * @Author: SpenserCai
 * @Date: 2023-08-14 17:36:46
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-09-08 15:14:59
 * @Description: file content
 */
package intersvc

// SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"

type LoraItemMetadata struct {
	SsSdModelName      string `json:"ss_sd_model_name"`
	SsBaseModelVersion string `json:"ss_base_model_version"`
}

type LoraItem struct {
	Name     string           `json:"name"`
	Alias    string           `json:"alias"`
	Path     string           `json:"path"`
	Metadata LoraItemMetadata `json:"metadata"`
}

type SdapiV1LorasResponse = []LoraItem
