package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /t2v/api_version
type T2vApiVersion struct {
	ResponseItem *T2vApiVersionResponse
	Error        error
}

func (d *T2vApiVersion) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewT2vAPIVersionT2vAPIVersionGetParams()
	ResponseData, err := inter.Client.Operations.T2vAPIVersionT2vAPIVersionGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &T2vApiVersionResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*T2vApiVersionResponse)
}

func (d *T2vApiVersion) GetResponse() *T2vApiVersionResponse {
	return d.ResponseItem
}
