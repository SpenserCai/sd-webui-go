package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /internal/progress
type InternalProgress struct {
	RequestItem  *InternalProgressRequest
	ResponseItem *InternalProgressResponse
	Error        error
}

func (d *InternalProgress) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewProgressapiInternalProgressPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.ProgressapiInternalProgressPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InternalProgressResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InternalProgressResponse)
}

func (d *InternalProgress) GetResponse() *InternalProgressResponse {
	return d.ResponseItem
}
