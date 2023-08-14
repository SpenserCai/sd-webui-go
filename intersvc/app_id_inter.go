package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /app_id
type AppId struct {
	ResponseItem *AppIdResponse
	Error        error
}

func (d *AppId) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewAppIDAppIDGetParams()
	ResponseData, err := inter.Client.Operations.AppIDAppIDGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &AppIdResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*AppIdResponse)
}

func (d *AppId) GetResponse() *AppIdResponse {
	return d.ResponseItem
}
