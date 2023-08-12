package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type SdapiV1Txt2img struct {
	RequestItem  *SdapiV1Txt2imgRequest
	ResponseItem *SdapiV1Txt2imgResponse
	Error        error
}

func (d *SdapiV1Txt2img) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewText2imgapiSdapiV1Txt2imgPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.Text2imgapiSdapiV1Txt2imgPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1Txt2imgResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1Txt2imgResponse)
}

func (d *SdapiV1Txt2img) GetResponse() *SdapiV1Txt2imgResponse {
	return d.ResponseItem
}