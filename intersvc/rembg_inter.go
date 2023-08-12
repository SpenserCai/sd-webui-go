package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type Rembg struct {
	RequestItem  *RembgRequest
	ResponseItem *RembgResponse
	Error        error
}

func (d *Rembg) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewRembgRemoveRembgPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.RembgRemoveRembgPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &RembgResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*RembgResponse)
}

func (d *Rembg) GetResponse() *RembgResponse {
	return d.ResponseItem
}