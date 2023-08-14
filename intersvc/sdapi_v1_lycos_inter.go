package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/lycos
type SdapiV1Lycos struct {
	ResponseItem *SdapiV1LycosResponse
	Error        error
}

func (d *SdapiV1Lycos) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetLycosSdapiV1LycosGetParams()
	ResponseData, err := inter.Client.Operations.GetLycosSdapiV1LycosGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1LycosResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1LycosResponse)
}

func (d *SdapiV1Lycos) GetResponse() *SdapiV1LycosResponse {
	return d.ResponseItem
}
