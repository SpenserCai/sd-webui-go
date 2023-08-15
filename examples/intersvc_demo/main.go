/*
 * @Author: SpenserCai
 * @Date: 2023-08-12 14:15:02
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-15 17:13:31
 * @Description: file content
 */
package main

import (
	"encoding/base64"
	"fmt"
	"os"

	SdClient "github.com/SpenserCai/sd-webui-go"
	"github.com/SpenserCai/sd-webui-go/intersvc"
)

func main() {
	// init client
	sdClient := SdClient.NewStableDiffInterface("127.0.0.1:7860")

	var f_factor int64 = 20
	var artistic bool = false

	// Set Request
	deoldify_inter := &intersvc.DeoldifyImage{
		RequestItem: &intersvc.DeoldifyImageRequest{
			InputImage:   "https://media.discordapp.net/attachments/1138408545277190237/1138508881635577947/i7krs1nj1ekla1.jpg",
			RenderFactor: &f_factor,
			Artistic:     &artistic,
		},
	}

	// Call Action
	deoldify_inter.Action(sdClient)
	if deoldify_inter.Error != nil {
		fmt.Println("Error: ", deoldify_inter.Error)
		return
	}

	response := deoldify_inter.GetResponse()

	reader, err := base64.StdEncoding.DecodeString(response.Image)
	if err != nil {
		panic(err)
	}
	os.WriteFile("output.jpg", reader, 0666)

}
