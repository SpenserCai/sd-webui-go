package intersvc

import (
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
		d.Error = err
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
