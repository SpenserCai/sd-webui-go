package intersvc

import (
	webui "github.com/SpenserCai/sd-webui-go"
	SdApiOperation "github.com/SpenserCai/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /sdapi/v1/embeddings
type SdapiV1Embeddings struct {
	ResponseItem *SdapiV1EmbeddingsResponse
	Error        error
}

func (d *SdapiV1Embeddings) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetEmbeddingsSdapiV1EmbeddingsGetParams()
	ResponseData, err := inter.Client.Operations.GetEmbeddingsSdapiV1EmbeddingsGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &SdapiV1EmbeddingsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*SdapiV1EmbeddingsResponse)
}

func (d *SdapiV1Embeddings) GetResponse() *SdapiV1EmbeddingsResponse {
	return d.ResponseItem
}
