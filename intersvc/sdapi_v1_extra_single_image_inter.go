package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type SdapiV1ExtraSingleImage struct {
	RequestItem  *SdapiV1ExtraSingleImageRequest
	ResponseItem *SdapiV1ExtraSingleImageResponse
	Error        error
}

func (d *SdapiV1ExtraSingleImage) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewExtrasSingleImageAPISdapiV1ExtraSingleImagePostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.ExtrasSingleImageAPISdapiV1ExtraSingleImagePost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1ExtraSingleImageResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1ExtraSingleImageResponse)
}

func (d *SdapiV1ExtraSingleImage) GetResponse() *SdapiV1ExtraSingleImageResponse {
	return d.ResponseItem
}