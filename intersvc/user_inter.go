package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /user
type User struct {
	ResponseItem *UserResponse
	Error        error
}

func (d *User) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetCurrentUserUserGetParams()
	ResponseData, err := inter.Client.Operations.GetCurrentUserUserGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &UserResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*UserResponse)
}

func (d *User) GetResponse() *UserResponse {
	return d.ResponseItem
}
