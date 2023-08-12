package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type SdapiV1ExtraBatchImages struct {
	RequestItem  *SdapiV1ExtraBatchImagesRequest
	ResponseItem *SdapiV1ExtraBatchImagesResponse
	Error        error
}

func (d *SdapiV1ExtraBatchImages) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1ExtraBatchImagesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1ExtraBatchImagesResponse)
}

func (d *SdapiV1ExtraBatchImages) GetResponse() *SdapiV1ExtraBatchImagesResponse {
	return d.ResponseItem
}