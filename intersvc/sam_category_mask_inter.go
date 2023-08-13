package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

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
		d.Error = err
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
