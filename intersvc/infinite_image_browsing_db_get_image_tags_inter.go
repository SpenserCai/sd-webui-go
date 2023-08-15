package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/get_image_tags
type InfiniteImageBrowsingDbGetImageTags struct {
	RequestItem  *InfiniteImageBrowsingDbGetImageTagsRequest
	ResponseItem *InfiniteImageBrowsingDbGetImageTagsResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbGetImageTags) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetImgTagsInfiniteImageBrowsingDbGetImageTagsPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.GetImgTagsInfiniteImageBrowsingDbGetImageTagsPost(RequestData)
	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(error(nil)) {
			d.Error = err
			return
		}
		errorValue := reflect.ValueOf(err).Elem().FieldByName("Payload")
		if !errorValue.IsValid() {
			d.Error = err
			return
		}
		d.Error = fmt.Errorf("%v", errorValue.Elem())
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbGetImageTagsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbGetImageTagsResponse)
}

func (d *InfiniteImageBrowsingDbGetImageTags) GetResponse() *InfiniteImageBrowsingDbGetImageTagsResponse {
	return d.ResponseItem
}
