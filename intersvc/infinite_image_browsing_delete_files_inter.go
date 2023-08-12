package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type InfiniteImageBrowsingDeleteFiles struct {
	RequestItem  *InfiniteImageBrowsingDeleteFilesRequest
	ResponseItem *InfiniteImageBrowsingDeleteFilesResponse
	Error        error
}

func (d *InfiniteImageBrowsingDeleteFiles) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewDeleteFilesInfiniteImageBrowsingDeleteFilesPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.DeleteFilesInfiniteImageBrowsingDeleteFilesPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDeleteFilesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDeleteFilesResponse)
}

func (d *InfiniteImageBrowsingDeleteFiles) GetResponse() *InfiniteImageBrowsingDeleteFilesResponse {
	return d.ResponseItem
}