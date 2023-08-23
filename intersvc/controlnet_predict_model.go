/*
 * @Author: SpenserCai
 * @Date: 2023-08-22 15:52:43
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-23 13:00:37
 * @Description: file content
 */
package intersvc

type ControlnetPredictArgsItem struct {
	Enabled       bool    `json:"enabled"`
	Module        string  `json:"module"`
	Model         string  `json:"model"`
	Weight        float64 `json:"weight"`
	Image         string  `json:"image"`
	ResizeMode    int64   `json:"resize_mode"`
	LowVRAM       bool    `json:"low_vram"`
	ProcessorRes  float64 `json:"processor_res"`
	ThresholdA    float64 `json:"threshold_a"`
	ThresholdB    float64 `json:"threshold_b"`
	GuidanceStart float64 `json:"guidance_start"`
	GuidanceEnd   float64 `json:"guidance_end"`
	ControlMode   int64   `json:"control_mode"`
	PixelPerFect  bool    `json:"pixel_perfect"`
}

type ControlnetPredictScript struct {
	Args []ControlnetPredictArgsItem `json:"args"`
}
