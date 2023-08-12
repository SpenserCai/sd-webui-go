package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type InfiniteImageBrowsingMoveFiles struct {
	RequestItem  *InfiniteImageBrowsingMoveFilesRequest
	ResponseItem *InfiniteImageBrowsingMoveFilesResponse
	Error        error
}

func (d *InfiniteImageBrowsingMoveFiles) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewMoveFilesInfiniteImageBrowsingMoveFilesPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.MoveFilesInfiniteImageBrowsingMoveFilesPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingMoveFilesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingMoveFilesResponse)
}

func (d *InfiniteImageBrowsingMoveFiles) GetResponse() *InfiniteImageBrowsingMoveFilesResponse {
	return d.ResponseItem
}