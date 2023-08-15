package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/extra-single-image
type SdapiV1ExtraSingleImage struct {
	RequestItem  *SdapiV1ExtraSingleImageRequest
	ResponseItem *SdapiV1ExtraSingleImageResponse
	Error        error
}

func (d *SdapiV1ExtraSingleImage) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewExtrasSingleImageAPISdapiV1ExtraSingleImagePostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.ExtrasSingleImageAPISdapiV1ExtraSingleImagePost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1ExtraSingleImageResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1ExtraSingleImageResponse)
}

func (d *SdapiV1ExtraSingleImage) GetResponse() *SdapiV1ExtraSingleImageResponse {
	return d.ResponseItem
}
