package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/hypernetworks
type SdapiV1Hypernetworks struct {
	ResponseItem *SdapiV1HypernetworksResponse
	Error        error
}

func (d *SdapiV1Hypernetworks) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetHypernetworksSdapiV1HypernetworksGetParams()
	ResponseData, err := inter.Client.Operations.GetHypernetworksSdapiV1HypernetworksGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1HypernetworksResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1HypernetworksResponse)
}

func (d *SdapiV1Hypernetworks) GetResponse() *SdapiV1HypernetworksResponse {
	return d.ResponseItem
}
