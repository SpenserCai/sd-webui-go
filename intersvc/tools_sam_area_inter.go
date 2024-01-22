package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /tools/sam_area
type ToolsSamArea struct {
	RequestItem  *ToolsSamAreaRequest
	ResponseItem *ToolsSamAreaResponse
	Error        error
}

func (d *ToolsSamArea) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewSamAreaToolsSamAreaPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.SamAreaToolsSamAreaPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &ToolsSamAreaResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*ToolsSamAreaResponse)
}

func (d *ToolsSamArea) GetResponse() *ToolsSamAreaResponse {
	return d.ResponseItem
}
