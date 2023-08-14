package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/hello
type InfiniteImageBrowsingHello struct {
	ResponseItem *InfiniteImageBrowsingHelloResponse
	Error        error
}

func (d *InfiniteImageBrowsingHello) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGreetingInfiniteImageBrowsingHelloGetParams()
	ResponseData, err := inter.Client.Operations.GreetingInfiniteImageBrowsingHelloGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingHelloResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingHelloResponse)
}

func (d *InfiniteImageBrowsingHello) GetResponse() *InfiniteImageBrowsingHelloResponse {
	return d.ResponseItem
}
