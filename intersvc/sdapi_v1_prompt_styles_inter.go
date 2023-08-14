package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/prompt-styles
type SdapiV1PromptStyles struct {
	ResponseItem *SdapiV1PromptStylesResponse
	Error        error
}

func (d *SdapiV1PromptStyles) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetPromptStylesSdapiV1PromptStylesGetParams()
	ResponseData, err := inter.Client.Operations.GetPromptStylesSdapiV1PromptStylesGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1PromptStylesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1PromptStylesResponse)
}

func (d *SdapiV1PromptStyles) GetResponse() *SdapiV1PromptStylesResponse {
	return d.ResponseItem
}
