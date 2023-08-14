package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/image-thumbnail
type InfiniteImageBrowsingImageThumbnail struct {
	ResponseItem *InfiniteImageBrowsingImageThumbnailResponse
	Error        error
}

func (d *InfiniteImageBrowsingImageThumbnail) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewThumbnailInfiniteImageBrowsingImageThumbnailGetParams()
	ResponseData, err := inter.Client.Operations.ThumbnailInfiniteImageBrowsingImageThumbnailGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingImageThumbnailResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingImageThumbnailResponse)
}

func (d *InfiniteImageBrowsingImageThumbnail) GetResponse() *InfiniteImageBrowsingImageThumbnailResponse {
	return d.ResponseItem
}
