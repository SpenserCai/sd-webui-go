package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /queue/status
type QueueStatus struct {
	ResponseItem *QueueStatusResponse
	Error        error
}

func (d *QueueStatus) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetQueueStatusQueueStatusGetParams()
	ResponseData, err := inter.Client.Operations.GetQueueStatusQueueStatusGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &QueueStatusResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*QueueStatusResponse)
}

func (d *QueueStatus) GetResponse() *QueueStatusResponse {
	return d.ResponseItem
}
