<!--
 * @Author: SpenserCai
 * @Date: 2023-08-12 01:27:12
 * @version: 
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-14 04:21:34
 * @Description: file content
-->

<div align="center">

<img src="https://raw.githubusercontent.com/SpenserCai/sd-webui-go/main/res/logo.png" width="200" height="200" alt="sd-webui-go">

# SD-WEBUI-GO
Golang SDK for [stable-diffusion-webui](https://github.com/AUTOMATIC1111/stable-diffusion-webui)'s API

<!--Among the contributors is AUTOMATIC1111, a little joke on my part.😂-->

<!--Here,sincerely thank AUTOMATIC1111 for its great contribution to AIGC-->

</div>

<p align="center">
  <a href="https://raw.githubusercontent.com/SpenserCai/sd-webui-go/main/LICENSE">
    <img src="https://img.shields.io/github/license/SpenserCai/sd-webui-go?color=blueviolet" alt="license">
  </a>
  <img src="https://img.shields.io/badge/Go-1.19+-blue" alt="go">
  <a href="https://github.com/SpenserCai/sd-webui-go/releases">
    <img src="https://img.shields.io/github/v/release/SpenserCai/sd-webui-go?color=rgb(255%2C0%2C0)&include_prereleases" alt="release">
  </a>
  <a href="https://goreportcard.com/report/github.com/SpenserCai/sd-webui-go">
    <img src="https://goreportcard.com/badge/github.com/SpenserCai/sd-webui-go" alt="GoReportCard">
  </a>
  <a href="https://discord.gg/nN6b7QAcVW">
  <img src="https://discordapp.com/api/guilds/1140177489008807966/widget.png?style=shield" alt="Discord Server">
</a>
</p>

This is a Go language version of the SDK based on [stable-diffusion-webui](https://github.com/AUTOMATIC1111/stable-diffusion-webui). In your code, you can directly use the API interfaces of [stable-diffusion-webui](https://github.com/AUTOMATIC1111/stable-diffusion-webui) through object-oriented operations, instead of dealing with cumbersome JSON.

Support extensions API !

<b>You can check the support list for intersvc in this [wiki page](https://github.com/SpenserCai/sd-webui-go/wiki/Intersvc-Support-List).</b>

## Usage
There are two methods to use the SDK. 

 - Using 'intersvc', which provides a highly encapsulated API interface for 'sd-webui'. Using this method feels like directly using the Go language version of 'sd-webui'. 

 - Using 'go-swagger', which still involves object-oriented operations but is slightly more complex than 'intersvc'. 

Almost all interfaces are supported by the second method, while the first one is gradually being supported.

In fact, most of the interfaces can be used with 'intersvc', but it requires defining the Response. The API documentation of 'sd-webui' does not provide the content of the Response, so it needs to be defined manually. You can refer to the "Participating" section below for specific instructions.


### intersvc

```go
import (
	SdClient "github.com/SpenserCai/sd-webui-go"
	"github.com/SpenserCai/sd-webui-go/intersvc"
)

func main() {
	// Create a client
	sdClient := SdClient.NewStableDiffInterface("127.0.0.1:7860")

	var f_factor int64 = 20
	var artistic bool = false

	// Set Request
	deoldify_inter := &intersvc.DeoldifyImage{
		RequestItem: &intersvc.DeoldifyImageRequest{
			InputImage:   "https://media.discordapp.net/attachments/1138408545277190237/1138508881635577947i7krs1njekla1.jpg",
			RenderFactor: &f_factor,
			Artistic:     &artistic,
		},
	}


	deoldify_inter.Action(sdClient)
	if deoldify_inter.Error != nil {
		panic(deoldify_inter.Error)
	}

	response := deoldify_inter.GetResponse()
}
```

Full example code: [intersvc_example](./examples/intersvc_demo/main.go)

### go-swagger

```go
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

    response := deoldifyRep.(*DeoldifyResponse)
}
```

Full example code: [go-swagger_example](./examples/stable_demo/main.go)

## Participating

Most of the code for intersvc has been generated using a code generator. However, due to the lack of response information in the API documentation for sd-webui, it needs to be manually written.

### How to submit a PR

You need to fork the code of the dev branch, make the necessary code updates, create a branch named dev-[model filename] in your own repository, and then submit a pull request to the dev branch of this repository.

### How to define the Response Model

In `sd-webui-go/intersvc`,you can see some file like `***_model.go`. These files define the Response Model. You can refer to the following example to define the Response Model.

From

```go
type DeoldifyImageResponse struct {
	
}
```

To

```go
type DeoldifyImageResponse struct {
	Image string `json:"image"`
}
```