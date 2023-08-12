package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type InternalProgress struct {
	RequestItem  *InternalProgressRequest
	ResponseItem *InternalProgressResponse
	Error        error
}

func (d *InternalProgress) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewProgressapiInternalProgressPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.ProgressapiInternalProgressPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InternalProgressResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InternalProgressResponse)
}

func (d *InternalProgress) GetResponse() *InternalProgressResponse {
	return d.ResponseItem
}