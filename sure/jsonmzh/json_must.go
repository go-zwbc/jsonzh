package jsonmzh

import (
	"github.com/go-zwbc/jsonzh"
	"github.com/yyle88/sure"
)

func M编码(object any) []byte {
	res0, err := jsonzh.M编码(object)
	sure.Must(err)
	return res0
}

func U解码[T any](data []byte) *T {
	res0, err := jsonzh.U解码[T](data)
	sure.Must(err)
	return res0
}

func M编码s(object any) string {
	res0, err := jsonzh.M编码s(object)
	sure.Must(err)
	return res0
}

func U解码s[T any](data string) *T {
	res0, err := jsonzh.U解码s[T](data)
	sure.Must(err)
	return res0
}
