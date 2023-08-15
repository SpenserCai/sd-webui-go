package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/match_images_by_tags
type InfiniteImageBrowsingDbMatchImagesByTags struct {
	RequestItem  *InfiniteImageBrowsingDbMatchImagesByTagsRequest
	ResponseItem *InfiniteImageBrowsingDbMatchImagesByTagsResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbMatchImagesByTags) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewMatchImageByTagsInfiniteImageBrowsingDbMatchImagesByTagsPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.MatchImageByTagsInfiniteImageBrowsingDbMatchImagesByTagsPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbMatchImagesByTagsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbMatchImagesByTagsResponse)
}

func (d *InfiniteImageBrowsingDbMatchImagesByTags) GetResponse() *InfiniteImageBrowsingDbMatchImagesByTagsResponse {
	return d.ResponseItem
}
