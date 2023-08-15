package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sam/dilate-mask
type SamDilateMask struct {
	RequestItem  *SamDilateMaskRequest
	ResponseItem *SamDilateMaskResponse
	Error        error
}

func (d *SamDilateMask) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewAPIDilateMaskSamDilateMaskPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.APIDilateMaskSamDilateMaskPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SamDilateMaskResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SamDilateMaskResponse)
}

func (d *SamDilateMask) GetResponse() *SamDilateMaskResponse {
	return d.ResponseItem
}
