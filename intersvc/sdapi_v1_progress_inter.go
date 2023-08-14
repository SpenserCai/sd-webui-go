package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/progress
type SdapiV1Progress struct {
	ResponseItem *SdapiV1ProgressResponse
	Error        error
}

func (d *SdapiV1Progress) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewProgressapiSdapiV1ProgressGetParams()
	ResponseData, err := inter.Client.Operations.ProgressapiSdapiV1ProgressGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1ProgressResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1ProgressResponse)
}

func (d *SdapiV1Progress) GetResponse() *SdapiV1ProgressResponse {
	return d.ResponseItem
}
