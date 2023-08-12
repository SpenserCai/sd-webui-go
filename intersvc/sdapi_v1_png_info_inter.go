package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type SdapiV1PngInfo struct {
	RequestItem  *SdapiV1PngInfoRequest
	ResponseItem *SdapiV1PngInfoResponse
	Error        error
}

func (d *SdapiV1PngInfo) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewPnginfoapiSdapiV1PngInfoPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.PnginfoapiSdapiV1PngInfoPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1PngInfoResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1PngInfoResponse)
}

func (d *SdapiV1PngInfo) GetResponse() *SdapiV1PngInfoResponse {
	return d.ResponseItem
}