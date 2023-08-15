package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sam/dino-predict
type SamDinoPredict struct {
	RequestItem  *SamDinoPredictRequest
	ResponseItem *SamDinoPredictResponse
	Error        error
}

func (d *SamDinoPredict) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewAPIDinoPredictSamDinoPredictPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.APIDinoPredictSamDinoPredictPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SamDinoPredictResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SamDinoPredictResponse)
}

func (d *SamDinoPredict) GetResponse() *SamDinoPredictResponse {
	return d.ResponseItem
}
