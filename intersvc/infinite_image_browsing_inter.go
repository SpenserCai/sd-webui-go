package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing
type InfiniteImageBrowsing struct {
	ResponseItem *InfiniteImageBrowsingResponse
	Error        error
}

func (d *InfiniteImageBrowsing) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewIndexBdInfiniteImageBrowsingGetParams()
	ResponseData, err := inter.Client.Operations.IndexBdInfiniteImageBrowsingGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingResponse)
}

func (d *InfiniteImageBrowsing) GetResponse() *InfiniteImageBrowsingResponse {
	return d.ResponseItem
}
