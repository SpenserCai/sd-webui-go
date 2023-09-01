/*
 * @Author: SpenserCai
 * @Date: 2023-08-11 13:15:25
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-09-01 13:15:08
 * @Description: file content
 */
package webui

import (
	"time"

	rclient "github.com/go-openapi/runtime/client"

	StableDiffClient "github.com/SpenserCai/sd-webui-go/stablediffusion/client"

	"github.com/go-openapi/strfmt"
)

type StableDiffInterface struct {
	Client *StableDiffClient.StableDiffusion
}

func NewStableDiffInterface(host string) *StableDiffInterface {
	rclient.DefaultTimeout = 600 * time.Second
	var client *StableDiffClient.StableDiffusion
	if host == "" {
		client = StableDiffClient.NewHTTPClient(strfmt.Default)
	} else {
		client = StableDiffClient.NewHTTPClientWithConfig(strfmt.Default, StableDiffClient.DefaultTransportConfig().WithHost(host))
	}
	return &StableDiffInterface{
		Client: client,
	}

}
