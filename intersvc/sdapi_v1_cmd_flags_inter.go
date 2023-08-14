package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/cmd-flags
type SdapiV1CmdFlags struct {
	ResponseItem *SdapiV1CmdFlagsResponse
	Error        error
}

func (d *SdapiV1CmdFlags) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetCmdFlagsSdapiV1CmdFlagsGetParams()
	ResponseData, err := inter.Client.Operations.GetCmdFlagsSdapiV1CmdFlagsGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1CmdFlagsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1CmdFlagsResponse)
}

func (d *SdapiV1CmdFlags) GetResponse() *SdapiV1CmdFlagsResponse {
	return d.ResponseItem
}
