package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sd_extra_networks/metadata
type SdExtraNetworksMetadata struct {
	ResponseItem *SdExtraNetworksMetadataResponse
	Error        error
}

func (d *SdExtraNetworksMetadata) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetMetadataSdExtraNetworksMetadataGetParams()
	ResponseData, err := inter.Client.Operations.GetMetadataSdExtraNetworksMetadataGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdExtraNetworksMetadataResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdExtraNetworksMetadataResponse)
}

func (d *SdExtraNetworksMetadata) GetResponse() *SdExtraNetworksMetadataResponse {
	return d.ResponseItem
}
