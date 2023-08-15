package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /rembg
type Rembg struct {
	RequestItem  *RembgRequest
	ResponseItem *RembgResponse
	Error        error
}

func (d *Rembg) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewRembgRemoveRembgPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.RembgRemoveRembgPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &RembgResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*RembgResponse)
}

func (d *Rembg) GetResponse() *RembgResponse {
	return d.ResponseItem
}
