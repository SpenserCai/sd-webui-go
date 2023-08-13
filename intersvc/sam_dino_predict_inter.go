package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

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
		d.Error = err
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
