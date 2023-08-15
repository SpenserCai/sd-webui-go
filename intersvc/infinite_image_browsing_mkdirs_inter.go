package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/mkdirs
type InfiniteImageBrowsingMkdirs struct {
	RequestItem  *InfiniteImageBrowsingMkdirsRequest
	ResponseItem *InfiniteImageBrowsingMkdirsResponse
	Error        error
}

func (d *InfiniteImageBrowsingMkdirs) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewCreateFoldersInfiniteImageBrowsingMkdirsPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.CreateFoldersInfiniteImageBrowsingMkdirsPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingMkdirsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingMkdirsResponse)
}

func (d *InfiniteImageBrowsingMkdirs) GetResponse() *InfiniteImageBrowsingMkdirsResponse {
	return d.ResponseItem
}
