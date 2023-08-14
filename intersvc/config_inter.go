package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /config
type Config struct {
	ResponseItem *ConfigResponse
	Error        error
}

func (d *Config) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetConfigConfigGetParams()
	ResponseData, err := inter.Client.Operations.GetConfigConfigGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &ConfigResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*ConfigResponse)
}

func (d *Config) GetResponse() *ConfigResponse {
	return d.ResponseItem
}
