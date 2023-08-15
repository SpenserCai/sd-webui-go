package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/delete_files
type InfiniteImageBrowsingDeleteFiles struct {
	RequestItem  *InfiniteImageBrowsingDeleteFilesRequest
	ResponseItem *InfiniteImageBrowsingDeleteFilesResponse
	Error        error
}

func (d *InfiniteImageBrowsingDeleteFiles) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewDeleteFilesInfiniteImageBrowsingDeleteFilesPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.DeleteFilesInfiniteImageBrowsingDeleteFilesPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDeleteFilesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDeleteFilesResponse)
}

func (d *InfiniteImageBrowsingDeleteFiles) GetResponse() *InfiniteImageBrowsingDeleteFilesResponse {
	return d.ResponseItem
}
