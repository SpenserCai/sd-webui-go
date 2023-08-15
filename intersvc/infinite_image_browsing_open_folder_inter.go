package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/open_folder
type InfiniteImageBrowsingOpenFolder struct {
	RequestItem  *InfiniteImageBrowsingOpenFolderRequest
	ResponseItem *InfiniteImageBrowsingOpenFolderResponse
	Error        error
}

func (d *InfiniteImageBrowsingOpenFolder) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewOpenFolderUsingExploreInfiniteImageBrowsingOpenFolderPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.OpenFolderUsingExploreInfiniteImageBrowsingOpenFolderPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingOpenFolderResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingOpenFolderResponse)
}

func (d *InfiniteImageBrowsingOpenFolder) GetResponse() *InfiniteImageBrowsingOpenFolderResponse {
	return d.ResponseItem
}
