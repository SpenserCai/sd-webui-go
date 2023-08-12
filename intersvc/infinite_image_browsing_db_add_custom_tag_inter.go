package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type InfiniteImageBrowsingDbAddCustomTag struct {
	RequestItem  *InfiniteImageBrowsingDbAddCustomTagRequest
	ResponseItem *InfiniteImageBrowsingDbAddCustomTagResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbAddCustomTag) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewAddCustomTagInfiniteImageBrowsingDbAddCustomTagPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.AddCustomTagInfiniteImageBrowsingDbAddCustomTagPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbAddCustomTagResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbAddCustomTagResponse)
}

func (d *InfiniteImageBrowsingDbAddCustomTag) GetResponse() *InfiniteImageBrowsingDbAddCustomTagResponse {
	return d.ResponseItem
}