package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sam/sam-model
type SamSamModel struct {
	ResponseItem *SamSamModelResponse
	Error        error
}

func (d *SamSamModel) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewAPISamModelSamSamModelGetParams()
	ResponseData, err := inter.Client.Operations.APISamModelSamSamModelGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SamSamModelResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SamSamModelResponse)
}

func (d *SamSamModel) GetResponse() *SamSamModelResponse {
	return d.ResponseItem
}
