package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /roop/image
type RoopImage struct {
	RequestItem  *RoopImageRequest
	ResponseItem *RoopImageResponse
	Error        error
}

func (d *RoopImage) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewRoopImageRoopImagePostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.RoopImageRoopImagePost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &RoopImageResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*RoopImageResponse)
}

func (d *RoopImage) GetResponse() *RoopImageResponse {
	return d.ResponseItem
}
