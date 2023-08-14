package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/loras
type SdapiV1Loras struct {
	ResponseItem *SdapiV1LorasResponse
	Error        error
}

func (d *SdapiV1Loras) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetLorasSdapiV1LorasGetParams()
	ResponseData, err := inter.Client.Operations.GetLorasSdapiV1LorasGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1LorasResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1LorasResponse)
}

func (d *SdapiV1Loras) GetResponse() *SdapiV1LorasResponse {
	return d.ResponseItem
}
