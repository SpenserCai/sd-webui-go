package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/txt2img
type SdapiV1Txt2img struct {
	RequestItem  *SdapiV1Txt2imgRequest
	ResponseItem *SdapiV1Txt2imgResponse
	Error        error
}

func (d *SdapiV1Txt2img) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewText2imgapiSdapiV1Txt2imgPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.Text2imgapiSdapiV1Txt2imgPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1Txt2imgResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1Txt2imgResponse)
}

func (d *SdapiV1Txt2img) GetResponse() *SdapiV1Txt2imgResponse {
	return d.ResponseItem
}
