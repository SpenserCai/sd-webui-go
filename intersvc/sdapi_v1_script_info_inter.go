package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/script-info
type SdapiV1ScriptInfo struct {
	ResponseItem *SdapiV1ScriptInfoResponse
	Error        error
}

func (d *SdapiV1ScriptInfo) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetScriptInfoSdapiV1ScriptInfoGetParams()
	ResponseData, err := inter.Client.Operations.GetScriptInfoSdapiV1ScriptInfoGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1ScriptInfoResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1ScriptInfoResponse)
}

func (d *SdapiV1ScriptInfo) GetResponse() *SdapiV1ScriptInfoResponse {
	return d.ResponseItem
}
