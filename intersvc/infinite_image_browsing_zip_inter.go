package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/zip
type InfiniteImageBrowsingZip struct {
	RequestItem  *InfiniteImageBrowsingZipRequest
	ResponseItem *InfiniteImageBrowsingZipResponse
	Error        error
}

func (d *InfiniteImageBrowsingZip) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewZipFilesInfiniteImageBrowsingZipPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.ZipFilesInfiniteImageBrowsingZipPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingZipResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingZipResponse)
}

func (d *InfiniteImageBrowsingZip) GetResponse() *InfiniteImageBrowsingZipResponse {
	return d.ResponseItem
}
