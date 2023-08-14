package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/gen_info_completed
type InfiniteImageBrowsingGenInfoCompleted struct {
	ResponseItem *InfiniteImageBrowsingGenInfoCompletedResponse
	Error        error
}

func (d *InfiniteImageBrowsingGenInfoCompleted) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewAPISetSendImgPathInfiniteImageBrowsingGenInfoCompletedGetParams()
	ResponseData, err := inter.Client.Operations.APISetSendImgPathInfiniteImageBrowsingGenInfoCompletedGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingGenInfoCompletedResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingGenInfoCompletedResponse)
}

func (d *InfiniteImageBrowsingGenInfoCompleted) GetResponse() *InfiniteImageBrowsingGenInfoCompletedResponse {
	return d.ResponseItem
}
