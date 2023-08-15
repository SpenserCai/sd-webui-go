package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sam/category-mask
type SamCategoryMask struct {
	RequestItem  *SamCategoryMaskRequest
	ResponseItem *SamCategoryMaskResponse
	Error        error
}

func (d *SamCategoryMask) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewAPICategoryMaskSamCategoryMaskPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.APICategoryMaskSamCategoryMaskPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SamCategoryMaskResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SamCategoryMaskResponse)
}

func (d *SamCategoryMask) GetResponse() *SamCategoryMaskResponse {
	return d.ResponseItem
}
