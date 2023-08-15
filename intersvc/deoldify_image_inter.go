package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /deoldify/image
type DeoldifyImage struct {
	RequestItem  *DeoldifyImageRequest
	ResponseItem *DeoldifyImageResponse
	Error        error
}

func (d *DeoldifyImage) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewDeoldifyImageDeoldifyImagePostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.DeoldifyImageDeoldifyImagePost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &DeoldifyImageResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*DeoldifyImageResponse)
}

func (d *DeoldifyImage) GetResponse() *DeoldifyImageResponse {
	return d.ResponseItem
}
