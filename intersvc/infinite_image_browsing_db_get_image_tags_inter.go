package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type InfiniteImageBrowsingDbGetImageTags struct {
	RequestItem  *InfiniteImageBrowsingDbGetImageTagsRequest
	ResponseItem *InfiniteImageBrowsingDbGetImageTagsResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbGetImageTags) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetImgTagsInfiniteImageBrowsingDbGetImageTagsPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.GetImgTagsInfiniteImageBrowsingDbGetImageTagsPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbGetImageTagsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbGetImageTagsResponse)
}

func (d *InfiniteImageBrowsingDbGetImageTags) GetResponse() *InfiniteImageBrowsingDbGetImageTagsResponse {
	return d.ResponseItem
}
