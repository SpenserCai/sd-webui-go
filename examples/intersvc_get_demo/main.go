/*
 * @Author: SpenserCai
 * @Date: 2023-08-14 21:51:22
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-14 22:23:31
 * @Description: file content
 */
package main

import (
	"fmt"

	SdClient "github.com/SpenserCai/sd-webui-go"
	"github.com/SpenserCai/sd-webui-go/intersvc"
)

func main() {
	// init client
	sdClient := SdClient.NewStableDiffInterface("127.0.0.1:7860")
	sd_models_inter := &intersvc.SdapiV1SdModels{}
	sd_models_inter.Action(sdClient)
	if sd_models_inter.Error != nil {
		panic(sd_models_inter.Error)
	}
	response := sd_models_inter.GetResponse()
	for _, item := range *response {
		// 遍历item的每个字段
		fmt.Println(*item.ModelName)
		fmt.Println(*item.Title)
		fmt.Println(*item.Filename)
		fmt.Println(item.Hash)
		fmt.Println(item.Sha256)
		fmt.Println(item.Config)
		fmt.Println("====================================")
	}

	sam_models_inter := &intersvc.SamSamModel{}
	sam_models_inter.Action(sdClient)
	if sam_models_inter.Error != nil {
		panic(sam_models_inter.Error)
	}
	sam_response := sam_models_inter.GetResponse()
	for _, item := range *sam_response {
		// 遍历item的每个字段
		fmt.Println(item)
	}

}
