package jsonzh

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func M编码(object any) ([]byte, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return nil, errors.WithMessage(err, "编码错误")
	}
	return data, nil
}

func U解码[T any](data []byte) (*T, error) {
	var object T
	err := json.Unmarshal(data, &object)
	if err != nil {
		return nil, errors.WithMessage(err, "解码错误") //当出错时直接暴露位置
	}
	return &object, nil
}

func M编码s(object any) (string, error) {
	data, err := M编码(object)
	if err != nil {
		return "", errors.WithMessage(err, "编码错误")
	}
	return string(data), nil
}

func U解码s[T any](data string) (*T, error) {
	object, err := U解码[T]([]byte(data))
	if err != nil {
		return nil, errors.WithMessage(err, "解码错误")
	}
	return object, nil
}
