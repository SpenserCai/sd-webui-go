package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/toggle_custom_tag_to_img
type InfiniteImageBrowsingDbToggleCustomTagToImg struct {
	RequestItem  *InfiniteImageBrowsingDbToggleCustomTagToImgRequest
	ResponseItem *InfiniteImageBrowsingDbToggleCustomTagToImgResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbToggleCustomTagToImg) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewToggleCustomTagToImgInfiniteImageBrowsingDbToggleCustomTagToImgPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.ToggleCustomTagToImgInfiniteImageBrowsingDbToggleCustomTagToImgPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbToggleCustomTagToImgResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbToggleCustomTagToImgResponse)
}

func (d *InfiniteImageBrowsingDbToggleCustomTagToImg) GetResponse() *InfiniteImageBrowsingDbToggleCustomTagToImgResponse {
	return d.ResponseItem
}
