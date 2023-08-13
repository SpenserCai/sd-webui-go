package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/open_folder
type InfiniteImageBrowsingOpenFolder struct {
	RequestItem  *InfiniteImageBrowsingOpenFolderRequest
	ResponseItem *InfiniteImageBrowsingOpenFolderResponse
	Error        error
}

func (d *InfiniteImageBrowsingOpenFolder) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewOpenFolderUsingExploreInfiniteImageBrowsingOpenFolderPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.OpenFolderUsingExploreInfiniteImageBrowsingOpenFolderPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingOpenFolderResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingOpenFolderResponse)
}

func (d *InfiniteImageBrowsingOpenFolder) GetResponse() *InfiniteImageBrowsingOpenFolderResponse {
	return d.ResponseItem
}
