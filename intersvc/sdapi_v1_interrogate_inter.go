package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

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
		d.Error = err
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