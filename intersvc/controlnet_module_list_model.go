/*
 * @Author: SpenserCai
 * @Date: 2023-08-14 17:36:46
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-20 04:11:40
 * @Description: file content
 */
package intersvc

// SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"

type ControlnetModuleDetailSlidersItem struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	Min   int    `json:"min"`
	Max   int    `json:"max"`
}

type ControlnetModuleDetailItem struct {
	ModelFree bool                                `json:"model_free"`
	Sliders   []ControlnetModuleDetailSlidersItem `json:"sliders"`
}

type ControlnetModuleListResponse struct {
	ModuleList   []string    `json:"module_list"`
	ModuleDetail interface{} `json:"module_detail"`
}
