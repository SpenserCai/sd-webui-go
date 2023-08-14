package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/search_by_substr
type InfiniteImageBrowsingDbSearchBySubstr struct {
	ResponseItem *InfiniteImageBrowsingDbSearchBySubstrResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbSearchBySubstr) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewSearchBySubstrInfiniteImageBrowsingDbSearchBySubstrGetParams()
	ResponseData, err := inter.Client.Operations.SearchBySubstrInfiniteImageBrowsingDbSearchBySubstrGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbSearchBySubstrResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbSearchBySubstrResponse)
}

func (d *InfiniteImageBrowsingDbSearchBySubstr) GetResponse() *InfiniteImageBrowsingDbSearchBySubstrResponse {
	return d.ResponseItem
}
