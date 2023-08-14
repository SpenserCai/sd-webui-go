package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sd_extra_networks/thumb
type SdExtraNetworksThumb struct {
	ResponseItem *SdExtraNetworksThumbResponse
	Error        error
}

func (d *SdExtraNetworksThumb) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewFetchFileSdExtraNetworksThumbGetParams()
	ResponseData, err := inter.Client.Operations.FetchFileSdExtraNetworksThumbGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdExtraNetworksThumbResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdExtraNetworksThumbResponse)
}

func (d *SdExtraNetworksThumb) GetResponse() *SdExtraNetworksThumbResponse {
	return d.ResponseItem
}
