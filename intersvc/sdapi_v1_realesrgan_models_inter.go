package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/realesrgan-models
type SdapiV1RealesrganModels struct {
	ResponseItem *SdapiV1RealesrganModelsResponse
	Error        error
}

func (d *SdapiV1RealesrganModels) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetRealesrganModelsSdapiV1RealesrganModelsGetParams()
	ResponseData, err := inter.Client.Operations.GetRealesrganModelsSdapiV1RealesrganModelsGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1RealesrganModelsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1RealesrganModelsResponse)
}

func (d *SdapiV1RealesrganModels) GetResponse() *SdapiV1RealesrganModelsResponse {
	return d.ResponseItem
}
