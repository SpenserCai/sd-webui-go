package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/global_setting
type InfiniteImageBrowsingGlobalSetting struct {
	ResponseItem *InfiniteImageBrowsingGlobalSettingResponse
	Error        error
}

func (d *InfiniteImageBrowsingGlobalSetting) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGlobalSettingInfiniteImageBrowsingGlobalSettingGetParams()
	ResponseData, err := inter.Client.Operations.GlobalSettingInfiniteImageBrowsingGlobalSettingGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingGlobalSettingResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingGlobalSettingResponse)
}

func (d *InfiniteImageBrowsingGlobalSetting) GetResponse() *InfiniteImageBrowsingGlobalSettingResponse {
	return d.ResponseItem
}
