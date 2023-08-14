package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/samplers
type SdapiV1Samplers struct {
	ResponseItem *SdapiV1SamplersResponse
	Error        error
}

func (d *SdapiV1Samplers) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetSamplersSdapiV1SamplersGetParams()
	ResponseData, err := inter.Client.Operations.GetSamplersSdapiV1SamplersGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1SamplersResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1SamplersResponse)
}

func (d *SdapiV1Samplers) GetResponse() *SdapiV1SamplersResponse {
	return d.ResponseItem
}
