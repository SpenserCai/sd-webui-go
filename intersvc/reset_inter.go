package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /reset
type Reset struct {
	RequestItem  *ResetRequest
	ResponseItem *ResetResponse
	Error        error
}

func (d *Reset) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewResetIteratorResetPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.ResetIteratorResetPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &ResetResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*ResetResponse)
}

func (d *Reset) GetResponse() *ResetResponse {
	return d.ResponseItem
}
