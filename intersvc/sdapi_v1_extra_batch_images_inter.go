package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/extra-batch-images
type SdapiV1ExtraBatchImages struct {
	RequestItem  *SdapiV1ExtraBatchImagesRequest
	ResponseItem *SdapiV1ExtraBatchImagesResponse
	Error        error
}

func (d *SdapiV1ExtraBatchImages) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1ExtraBatchImagesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1ExtraBatchImagesResponse)
}

func (d *SdapiV1ExtraBatchImages) GetResponse() *SdapiV1ExtraBatchImagesResponse {
	return d.ResponseItem
}
