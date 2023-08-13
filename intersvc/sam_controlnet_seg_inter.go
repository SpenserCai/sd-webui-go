package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sam/controlnet-seg
type SamControlnetSeg struct {
	RequestItem  *SamControlnetSegRequest
	ResponseItem *SamControlnetSegResponse
	Error        error
}

func (d *SamControlnetSeg) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewAPIControlnetSegSamControlnetSegPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.APIControlnetSegSamControlnetSegPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SamControlnetSegResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SamControlnetSegResponse)
}

func (d *SamControlnetSeg) GetResponse() *SamControlnetSegResponse {
	return d.ResponseItem
}
