package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /t2v/version
type T2vVersion struct {
	ResponseItem *T2vVersionResponse
	Error        error
}

func (d *T2vVersion) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewT2vVersionT2vVersionGetParams()
	ResponseData, err := inter.Client.Operations.T2vVersionT2vVersionGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &T2vVersionResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*T2vVersionResponse)
}

func (d *T2vVersion) GetResponse() *T2vVersionResponse {
	return d.ResponseItem
}
