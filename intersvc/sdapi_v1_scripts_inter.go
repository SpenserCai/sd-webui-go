package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/scripts
type SdapiV1Scripts struct {
	ResponseItem *SdapiV1ScriptsResponse
	Error        error
}

func (d *SdapiV1Scripts) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetScriptsListSdapiV1ScriptsGetParams()
	ResponseData, err := inter.Client.Operations.GetScriptsListSdapiV1ScriptsGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1ScriptsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1ScriptsResponse)
}

func (d *SdapiV1Scripts) GetResponse() *SdapiV1ScriptsResponse {
	return d.ResponseItem
}
