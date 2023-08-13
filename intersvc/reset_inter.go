package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /reset
type Reset struct {
	RequestItem  *ResetRequest
	ResponseItem *ResetResponse
	Error        error
}

func (d *Reset) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewResetIteratorResetPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.ResetIteratorResetPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &ResetResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*ResetResponse)
}

func (d *Reset) GetResponse() *ResetResponse {
	return d.ResponseItem
}
