<!--
 * @Author: SpenserCai
 * @Date: 2023-08-12 01:27:12
 * @version: 
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-12 20:01:49
 * @Description: file content
-->
# SD-WEBUI-GO
This is a Go language version of the SDK based on [stable-diffusion-webui](https://github.com/AUTOMATIC1111/stable-diffusion-webui). In your code, you can directly use the API interfaces of [stable-diffusion-webui](https://github.com/AUTOMATIC1111/stable-diffusion-webui) through object-oriented operations, instead of dealing with cumbersome JSON.

Support extensions API !

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
			InputImage:   "https://media.discordapp.net/attachments/       1138408545277190237/1138508881635577947/i7krs1njekla1.       jpg",
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