package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/remove_custom_tag
type InfiniteImageBrowsingDbRemoveCustomTag struct {
	RequestItem  *InfiniteImageBrowsingDbRemoveCustomTagRequest
	ResponseItem *InfiniteImageBrowsingDbRemoveCustomTagResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbRemoveCustomTag) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewRemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbRemoveCustomTagResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbRemoveCustomTagResponse)
}

func (d *InfiniteImageBrowsingDbRemoveCustomTag) GetResponse() *InfiniteImageBrowsingDbRemoveCustomTagResponse {
	return d.ResponseItem
}
