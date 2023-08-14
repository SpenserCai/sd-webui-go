package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/files
type InfiniteImageBrowsingFiles struct {
	ResponseItem *InfiniteImageBrowsingFilesResponse
	Error        error
}

func (d *InfiniteImageBrowsingFiles) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetTargetFloderFilesInfiniteImageBrowsingFilesGetParams()
	ResponseData, err := inter.Client.Operations.GetTargetFloderFilesInfiniteImageBrowsingFilesGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingFilesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingFilesResponse)
}

func (d *InfiniteImageBrowsingFiles) GetResponse() *InfiniteImageBrowsingFilesResponse {
	return d.ResponseItem
}
