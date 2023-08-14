package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/upscalers
type SdapiV1Upscalers struct {
	ResponseItem *SdapiV1UpscalersResponse
	Error        error
}

func (d *SdapiV1Upscalers) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetUpscalersSdapiV1UpscalersGetParams()
	ResponseData, err := inter.Client.Operations.GetUpscalersSdapiV1UpscalersGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1UpscalersResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1UpscalersResponse)
}

func (d *SdapiV1Upscalers) GetResponse() *SdapiV1UpscalersResponse {
	return d.ResponseItem
}
