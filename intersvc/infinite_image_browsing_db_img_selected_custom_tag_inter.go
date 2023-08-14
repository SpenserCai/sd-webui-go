package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/img_selected_custom_tag
type InfiniteImageBrowsingDbImgSelectedCustomTag struct {
	ResponseItem *InfiniteImageBrowsingDbImgSelectedCustomTagResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbImgSelectedCustomTag) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetImgSelectedCustomTagInfiniteImageBrowsingDbImgSelectedCustomTagGetParams()
	ResponseData, err := inter.Client.Operations.GetImgSelectedCustomTagInfiniteImageBrowsingDbImgSelectedCustomTagGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbImgSelectedCustomTagResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbImgSelectedCustomTagResponse)
}

func (d *InfiniteImageBrowsingDbImgSelectedCustomTag) GetResponse() *InfiniteImageBrowsingDbImgSelectedCustomTagResponse {
	return d.ResponseItem
}
