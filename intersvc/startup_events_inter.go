package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /startup-events
type StartupEvents struct {
	ResponseItem *StartupEventsResponse
	Error        error
}

func (d *StartupEvents) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewStartupEventsStartupEventsGetParams()
	ResponseData, err := inter.Client.Operations.StartupEventsStartupEventsGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &StartupEventsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*StartupEventsResponse)
}

func (d *StartupEvents) GetResponse() *StartupEventsResponse {
	return d.ResponseItem
}
