package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /login_check
type LoginCheck struct {
	ResponseItem *LoginCheckResponse
	Error        error
}

func (d *LoginCheck) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewLoginCheckLoginCheckGetParams()
	ResponseData, err := inter.Client.Operations.LoginCheckLoginCheckGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &LoginCheckResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*LoginCheckResponse)
}

func (d *LoginCheck) GetResponse() *LoginCheckResponse {
	return d.ResponseItem
}
