package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/image_geninfo
type InfiniteImageBrowsingImageGeninfo struct {
	ResponseItem *InfiniteImageBrowsingImageGeninfoResponse
	Error        error
}

func (d *InfiniteImageBrowsingImageGeninfo) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewImageGeninfoInfiniteImageBrowsingImageGeninfoGetParams()
	ResponseData, err := inter.Client.Operations.ImageGeninfoInfiniteImageBrowsingImageGeninfoGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingImageGeninfoResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingImageGeninfoResponse)
}

func (d *InfiniteImageBrowsingImageGeninfo) GetResponse() *InfiniteImageBrowsingImageGeninfoResponse {
	return d.ResponseItem
}
