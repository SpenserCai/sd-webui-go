package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/remove_custom_tag_from_img
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
