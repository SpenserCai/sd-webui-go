package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /controlnet/module_list
type ControlnetModuleList struct {
	ResponseItem *ControlnetModuleListResponse
	Error        error
}

func (d *ControlnetModuleList) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewModuleListControlnetModuleListGetParams()
	ResponseData, err := inter.Client.Operations.ModuleListControlnetModuleListGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &ControlnetModuleListResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*ControlnetModuleListResponse)
}

func (d *ControlnetModuleList) GetResponse() *ControlnetModuleListResponse {
	return d.ResponseItem
}
