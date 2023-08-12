/*
 * @Author: SpenserCai
 * @Date: 2023-08-11 22:37:37
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-12 00:53:51
 * @Description: file content
 */
package intersvc

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 从指定类型转换另一个类型，第一个参数是interface{}，第二个参数是类型
func ConvertResponse(response interface{}, target interface{}) (interface{}, error) {
	targetType := reflect.TypeOf(target)
	newValuePtr := reflect.New(targetType) // 创建指向目标类型的指针
	newValue := newValuePtr.Elem()         // 获取指针指向的值

	// 将 response 转换为 JSON 字节
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %v", err)
	}

	// 将 JSON 字节解码到新的目标类型变量中
	if err := json.Unmarshal(responseBytes, newValue.Addr().Interface()); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return newValue.Interface(), nil
}
