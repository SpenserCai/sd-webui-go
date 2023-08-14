package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/memory
type SdapiV1Memory struct {
	ResponseItem *SdapiV1MemoryResponse
	Error        error
}

func (d *SdapiV1Memory) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetMemorySdapiV1MemoryGetParams()
	ResponseData, err := inter.Client.Operations.GetMemorySdapiV1MemoryGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1MemoryResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1MemoryResponse)
}

func (d *SdapiV1Memory) GetResponse() *SdapiV1MemoryResponse {
	return d.ResponseItem
}
