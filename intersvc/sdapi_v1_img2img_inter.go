package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/img2img
type SdapiV1Img2img struct {
	RequestItem  *SdapiV1Img2imgRequest
	ResponseItem *SdapiV1Img2imgResponse
	Error        error
}

func (d *SdapiV1Img2img) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewImg2imgapiSdapiV1Img2imgPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.Img2imgapiSdapiV1Img2imgPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1Img2imgResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1Img2imgResponse)
}

func (d *SdapiV1Img2img) GetResponse() *SdapiV1Img2imgResponse {
	return d.ResponseItem
}
