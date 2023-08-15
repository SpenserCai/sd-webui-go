package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/interrogate
type SdapiV1Interrogate struct {
	RequestItem  *SdapiV1InterrogateRequest
	ResponseItem *SdapiV1InterrogateResponse
	Error        error
}

func (d *SdapiV1Interrogate) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewInterrogateapiSdapiV1InterrogatePostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.InterrogateapiSdapiV1InterrogatePost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1InterrogateResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1InterrogateResponse)
}

func (d *SdapiV1Interrogate) GetResponse() *SdapiV1InterrogateResponse {
	return d.ResponseItem
}
