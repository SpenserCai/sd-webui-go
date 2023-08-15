package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/move_files
type InfiniteImageBrowsingMoveFiles struct {
	RequestItem  *InfiniteImageBrowsingMoveFilesRequest
	ResponseItem *InfiniteImageBrowsingMoveFilesResponse
	Error        error
}

func (d *InfiniteImageBrowsingMoveFiles) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewMoveFilesInfiniteImageBrowsingMoveFilesPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.MoveFilesInfiniteImageBrowsingMoveFilesPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingMoveFilesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingMoveFilesResponse)
}

func (d *InfiniteImageBrowsingMoveFiles) GetResponse() *InfiniteImageBrowsingMoveFilesResponse {
	return d.ResponseItem
}
