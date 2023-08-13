package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type InfiniteImageBrowsingCheckPathExists struct {
	RequestItem  *InfiniteImageBrowsingCheckPathExistsRequest
	ResponseItem *InfiniteImageBrowsingCheckPathExistsResponse
	Error        error
}

func (d *InfiniteImageBrowsingCheckPathExists) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewCheckPathExistsInfiniteImageBrowsingCheckPathExistsPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.CheckPathExistsInfiniteImageBrowsingCheckPathExistsPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingCheckPathExistsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingCheckPathExistsResponse)
}

func (d *InfiniteImageBrowsingCheckPathExists) GetResponse() *InfiniteImageBrowsingCheckPathExistsResponse {
	return d.ResponseItem
}
