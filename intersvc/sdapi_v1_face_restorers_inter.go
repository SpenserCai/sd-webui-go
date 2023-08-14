package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/face-restorers
type SdapiV1FaceRestorers struct {
	ResponseItem *SdapiV1FaceRestorersResponse
	Error        error
}

func (d *SdapiV1FaceRestorers) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetFaceRestorersSdapiV1FaceRestorersGetParams()
	ResponseData, err := inter.Client.Operations.GetFaceRestorersSdapiV1FaceRestorersGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1FaceRestorersResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1FaceRestorersResponse)
}

func (d *SdapiV1FaceRestorers) GetResponse() *SdapiV1FaceRestorersResponse {
	return d.ResponseItem
}
