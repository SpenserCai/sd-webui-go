package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sam/heartbeat
type SamHeartbeat struct {
	ResponseItem *SamHeartbeatResponse
	Error        error
}

func (d *SamHeartbeat) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewHeartbeatSamHeartbeatGetParams()
	ResponseData, err := inter.Client.Operations.HeartbeatSamHeartbeatGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SamHeartbeatResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SamHeartbeatResponse)
}

func (d *SamHeartbeat) GetResponse() *SamHeartbeatResponse {
	return d.ResponseItem
}
