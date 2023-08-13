package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type InfiniteImageBrowsingCopyFiles struct {
	RequestItem  *InfiniteImageBrowsingCopyFilesRequest
	ResponseItem *InfiniteImageBrowsingCopyFilesResponse
	Error        error
}

func (d *InfiniteImageBrowsingCopyFiles) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewCopyFilesInfiniteImageBrowsingCopyFilesPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.CopyFilesInfiniteImageBrowsingCopyFilesPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingCopyFilesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingCopyFilesResponse)
}

func (d *InfiniteImageBrowsingCopyFiles) GetResponse() *InfiniteImageBrowsingCopyFilesResponse {
	return d.ResponseItem
}
