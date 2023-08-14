package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/sd-models
type SdapiV1SdModels struct {
	ResponseItem *SdapiV1SdModelsResponse
	Error        error
}

func (d *SdapiV1SdModels) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetSdModelsSdapiV1SdModelsGetParams()
	ResponseData, err := inter.Client.Operations.GetSdModelsSdapiV1SdModelsGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1SdModelsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1SdModelsResponse)
}

func (d *SdapiV1SdModels) GetResponse() *SdapiV1SdModelsResponse {
	return d.ResponseItem
}
