package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/sd-vae
type SdapiV1SdVae struct {
	ResponseItem *SdapiV1SdVaeResponse
	Error        error
}

func (d *SdapiV1SdVae) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetSdVaesSdapiV1SdVaeGetParams()
	ResponseData, err := inter.Client.Operations.GetSdVaesSdapiV1SdVaeGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1SdVaeResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1SdVaeResponse)
}

func (d *SdapiV1SdVae) GetResponse() *SdapiV1SdVaeResponse {
	return d.ResponseItem
}
