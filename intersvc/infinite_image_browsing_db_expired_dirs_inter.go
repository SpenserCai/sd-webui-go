package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/expired_dirs
type InfiniteImageBrowsingDbExpiredDirs struct {
	ResponseItem *InfiniteImageBrowsingDbExpiredDirsResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbExpiredDirs) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetDbExpiredInfiniteImageBrowsingDbExpiredDirsGetParams()
	ResponseData, err := inter.Client.Operations.GetDbExpiredInfiniteImageBrowsingDbExpiredDirsGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbExpiredDirsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbExpiredDirsResponse)
}

func (d *InfiniteImageBrowsingDbExpiredDirs) GetResponse() *InfiniteImageBrowsingDbExpiredDirsResponse {
	return d.ResponseItem
}
