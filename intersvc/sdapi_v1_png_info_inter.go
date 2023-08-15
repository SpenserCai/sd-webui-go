package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/png-info
type SdapiV1PngInfo struct {
	RequestItem  *SdapiV1PngInfoRequest
	ResponseItem *SdapiV1PngInfoResponse
	Error        error
}

func (d *SdapiV1PngInfo) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewPnginfoapiSdapiV1PngInfoPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.PnginfoapiSdapiV1PngInfoPost(RequestData)
	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(error(nil)) {
			d.Error = err
			return
		}
		errorValue := reflect.ValueOf(err).Elem().FieldByName("Payload")
		if !errorValue.IsValid() {
			d.Error = err
			return
		}
		d.Error = fmt.Errorf("%v", errorValue.Elem())
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1PngInfoResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1PngInfoResponse)
}

func (d *SdapiV1PngInfo) GetResponse() *SdapiV1PngInfoResponse {
	return d.ResponseItem
}
