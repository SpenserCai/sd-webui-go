package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type SamSamPredict struct {
	RequestItem  *SamSamPredictRequest
	ResponseItem *SamSamPredictResponse
	Error        error
}

func (d *SamSamPredict) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewAPISamPredictSamSamPredictPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.APISamPredictSamSamPredictPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SamSamPredictResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SamSamPredictResponse)
}

func (d *SamSamPredict) GetResponse() *SamSamPredictResponse {
	return d.ResponseItem
}
