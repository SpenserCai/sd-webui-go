/*
 * @Author: SpenserCai
 * @Date: 2023-08-12 00:43:12
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-12 15:23:00
 * @Description: file content
 */
package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

type DeoldifyImage struct {
	RequestItem  *DeoldifyImageRequest
	ResponseItem *DeoldifyImageResponse
	Error        error
}

func (d *DeoldifyImage) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewDeoldifyImageDeoldifyImagePostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.DeoldifyImageDeoldifyImagePost(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &DeoldifyImageResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*DeoldifyImageResponse)
}

func (d *DeoldifyImage) GetResponse() *DeoldifyImageResponse {
	return d.ResponseItem
}
