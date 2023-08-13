package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type InfiniteImageBrowsingDbRemoveCustomTagFromImg struct {
	RequestItem  *InfiniteImageBrowsingDbRemoveCustomTagFromImgRequest
	ResponseItem *InfiniteImageBrowsingDbRemoveCustomTagFromImgResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbRemoveCustomTagFromImg) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewRemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbRemoveCustomTagFromImgResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbRemoveCustomTagFromImgResponse)
}

func (d *InfiniteImageBrowsingDbRemoveCustomTagFromImg) GetResponse() *InfiniteImageBrowsingDbRemoveCustomTagFromImgResponse {
	return d.ResponseItem
}
