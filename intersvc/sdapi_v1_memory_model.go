/*
 * @Author: SpenserCai
 * @Date: 2023-08-14 17:36:46
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-21 18:56:30
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type MemoryRAMInfo struct {
	Free  int64 `json:"free"`
	Used  int64 `json:"used"`
	Total int64 `json:"total"`
}

type MemoryCudaSystemInfo struct {
	Free  int64 `json:"free"`
	Used  int64 `json:"used"`
	Total int64 `json:"total"`
}

type MemoryCudaInfo struct {
	System MemoryCudaSystemInfo `json:"system"`
}

type SdapiV1MemoryResponse = SdApiModel.MemoryResponse
