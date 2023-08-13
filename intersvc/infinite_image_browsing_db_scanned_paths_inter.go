package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type InfiniteImageBrowsingDbScannedPaths struct {
	RequestItem  *InfiniteImageBrowsingDbScannedPathsRequest
	ResponseItem *InfiniteImageBrowsingDbScannedPathsResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbScannedPaths) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewCreateScannedPathInfiniteImageBrowsingDbScannedPathsPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.CreateScannedPathInfiniteImageBrowsingDbScannedPathsPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbScannedPathsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbScannedPathsResponse)
}

func (d *InfiniteImageBrowsingDbScannedPaths) GetResponse() *InfiniteImageBrowsingDbScannedPathsResponse {
	return d.ResponseItem
}
