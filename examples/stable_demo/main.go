/*
 * @Author: SpenserCai
 * @Date: 2023-08-12 14:16:04
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-12 15:03:10
 * @Description: file content
 */
package main

import (
	"encoding/base64"
	"os"

	SdClient "github.com/SpenserCai/sd-webui-go"
	intersvc "github.com/SpenserCai/sd-webui-go/intersvc"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type DeoldifyResponse struct {
	Image string `json:"image"`
}

func main() {
	var (
		err error
	)

	// init client
	sdClient := SdClient.NewStableDiffInterface("127.0.0.1:7860")

	var f_factor int64 = 20
	var artistic bool = false

	// set request data
	RequestData := SdApiOperation.NewDeoldifyImageDeoldifyImagePostParams()
	RequestData.Body = &SdApiModel.BodyDeoldifyImageDeoldifyImagePost{
		InputImage:   "https://media.discordapp.net/attachments/1138408545277190237/1138508881635577947/i7krs1njekla1.jpg",
		RenderFactor: &f_factor,
		Artistic:     &artistic,
	}

	// send request
	Response, err := sdClient.Client.Operations.DeoldifyImageDeoldifyImagePost(RequestData)
	if err != nil {
		panic(err)
	}

	// convert response
	deoldifyRep, err := intersvc.ConvertResponse(Response.Payload, &DeoldifyResponse{})
	if err != nil {
		panic(err)
	}

	reader, err := base64.StdEncoding.DecodeString(deoldifyRep.(*DeoldifyResponse).Image)
	if err != nil {
		panic(err)
	}

	os.WriteFile("output.jpg", reader, 0666)

}
