package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/add_custom_tag
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
