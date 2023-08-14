package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /controlnet/version
type ControlnetVersion struct {
	ResponseItem *ControlnetVersionResponse
	Error        error
}

func (d *ControlnetVersion) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewVersionControlnetVersionGetParams()
	ResponseData, err := inter.Client.Operations.VersionControlnetVersionGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &ControlnetVersionResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*ControlnetVersionResponse)
}

func (d *ControlnetVersion) GetResponse() *ControlnetVersionResponse {
	return d.ResponseItem
}
