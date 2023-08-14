package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/latent-upscale-modes
type SdapiV1LatentUpscaleModes struct {
	ResponseItem *SdapiV1LatentUpscaleModesResponse
	Error        error
}

func (d *SdapiV1LatentUpscaleModes) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetLatentUpscaleModesSdapiV1LatentUpscaleModesGetParams()
	ResponseData, err := inter.Client.Operations.GetLatentUpscaleModesSdapiV1LatentUpscaleModesGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1LatentUpscaleModesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1LatentUpscaleModesResponse)
}

func (d *SdapiV1LatentUpscaleModes) GetResponse() *SdapiV1LatentUpscaleModesResponse {
	return d.ResponseItem
}
