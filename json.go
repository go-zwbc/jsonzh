// Package jsonzh provides JSON encoding/decoding with Chinese function names
// Wraps Go's standard encoding/json package with enhanced error context
// Features type-safe generic operations and both byte/string format support
// Includes multiple error handling strategies via subpackages
//
// jsonzh 提供使用中文函数名的 JSON 编解码功能
// 封装 Go 标准 encoding/json 包，增强错误上下文信息
// 支持类型安全的泛型操作，同时支持字节和字符串格式
// 通过子包提供多种错误处理策略
package jsonzh

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// M编码 encodes any object into JSON bytes with enhanced error context
// object: the Go value to encode
// Returns JSON bytes and error with stack trace when encoding fails
//
// M编码 将任意对象编码为 JSON 字节，提供增强的错误上下文
// object: 要编码的 Go 值
// 返回 JSON 字节，编码失败时返回带堆栈跟踪的错误
func M编码(object any) ([]byte, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return nil, errors.WithMessage(err, "编码错误")
	}
	return data, nil
}

// U解码 decodes JSON bytes into a typed object with enhanced error context
// T: the target type to decode into
// data: the JSON bytes to decode
// Returns decoded object as pointer, plus error with stack trace when decode fails
//
// U解码 将 JSON 字节解码为类型化对象，提供增强的错误上下文
// T: 要解码到的目标类型
// data: 要解码的 JSON 字节
// 返回解码后对象的指针，解码失败时返回带堆栈跟踪的错误
func U解码[T any](data []byte) (*T, error) {
	var object T
	err := json.Unmarshal(data, &object)
	if err != nil {
		return nil, errors.WithMessage(err, "解码错误") //当出错时直接暴露位置
	}
	return &object, nil
}

// M编码s encodes any object into a JSON string with enhanced error context
// object: the Go value to encode
// Returns JSON string and error with stack trace when encoding fails
//
// M编码s 将任意对象编码为 JSON 字符串，提供增强的错误上下文
// object: 要编码的 Go 值
// 返回 JSON 字符串，编码失败时返回带堆栈跟踪的错误
func M编码s(object any) (string, error) {
	data, err := M编码(object)
	if err != nil {
		return "", errors.WithMessage(err, "编码错误")
	}
	return string(data), nil
}

// U解码s decodes a JSON string into a typed object with enhanced error context
// T: the target type to decode into
// data: the JSON string to decode
// Returns decoded object as pointer, plus error with stack trace when decode fails
//
// U解码s 将 JSON 字符串解码为类型化对象，提供增强的错误上下文
// T: 要解码到的目标类型
// data: 要解码的 JSON 字符串
// 返回解码后对象的指针，解码失败时返回带堆栈跟踪的错误
func U解码s[T any](data string) (*T, error) {
	object, err := U解码[T]([]byte(data))
	if err != nil {
		return nil, errors.WithMessage(err, "解码错误")
	}
	return object, nil
}
