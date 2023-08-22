package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /roop/models
type RoopModels struct {
	ResponseItem *RoopModelsResponse
	Error        error
}

func (d *RoopModels) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewRoopModelsRoopModelsGetParams()
	ResponseData, err := inter.Client.Operations.RoopModelsRoopModelsGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &RoopModelsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*RoopModelsResponse)
}

func (d *RoopModels) GetResponse() *RoopModelsResponse {
	return d.ResponseItem
}
