package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /controlnet/settings
type ControlnetSettings struct {
	ResponseItem *ControlnetSettingsResponse
	Error        error
}

func (d *ControlnetSettings) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewSettingsControlnetSettingsGetParams()
	ResponseData, err := inter.Client.Operations.SettingsControlnetSettingsGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &ControlnetSettingsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*ControlnetSettingsResponse)
}

func (d *ControlnetSettings) GetResponse() *ControlnetSettingsResponse {
	return d.ResponseItem
}
