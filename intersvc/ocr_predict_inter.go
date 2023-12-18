package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /ocr/predict
type OcrPredict struct {
	RequestItem  *OcrPredictRequest
	ResponseItem *OcrPredictResponse
	Error        error
}

func (d *OcrPredict) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewPredictOcrPredictPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.PredictOcrPredictPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &OcrPredictResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*OcrPredictResponse)
}

func (d *OcrPredict) GetResponse() *OcrPredictResponse {
	return d.ResponseItem
}
