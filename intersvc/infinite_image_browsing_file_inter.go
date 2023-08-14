package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/file
type InfiniteImageBrowsingFile struct {
	ResponseItem *InfiniteImageBrowsingFileResponse
	Error        error
}

func (d *InfiniteImageBrowsingFile) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetFileInfiniteImageBrowsingFileGetParams()
	ResponseData, err := inter.Client.Operations.GetFileInfiniteImageBrowsingFileGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingFileResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingFileResponse)
}

func (d *InfiniteImageBrowsingFile) GetResponse() *InfiniteImageBrowsingFileResponse {
	return d.ResponseItem
}
