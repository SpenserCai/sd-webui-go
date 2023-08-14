package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sd_extra_networks/get-single-card
type SdExtraNetworksGetSingleCard struct {
	ResponseItem *SdExtraNetworksGetSingleCardResponse
	Error        error
}

func (d *SdExtraNetworksGetSingleCard) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetSingleCardSdExtraNetworksGetSingleCardGetParams()
	ResponseData, err := inter.Client.Operations.GetSingleCardSdExtraNetworksGetSingleCardGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdExtraNetworksGetSingleCardResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdExtraNetworksGetSingleCardResponse)
}

func (d *SdExtraNetworksGetSingleCard) GetResponse() *SdExtraNetworksGetSingleCardResponse {
	return d.ResponseItem
}
