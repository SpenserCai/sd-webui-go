package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /controlnet/model_list
type ControlnetModelList struct {
	ResponseItem *ControlnetModelListResponse
	Error        error
}

func (d *ControlnetModelList) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewModelListControlnetModelListGetParams()
	ResponseData, err := inter.Client.Operations.ModelListControlnetModelListGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &ControlnetModelListResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*ControlnetModelListResponse)
}

func (d *ControlnetModelList) GetResponse() *ControlnetModelListResponse {
	return d.ResponseItem
}
