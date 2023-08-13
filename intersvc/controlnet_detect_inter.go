package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /controlnet/detect
type ControlnetDetect struct {
	RequestItem  *ControlnetDetectRequest
	ResponseItem *ControlnetDetectResponse
	Error        error
}

func (d *ControlnetDetect) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewDetectControlnetDetectPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.DetectControlnetDetectPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &ControlnetDetectResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*ControlnetDetectResponse)
}

func (d *ControlnetDetect) GetResponse() *ControlnetDetectResponse {
	return d.ResponseItem
}
