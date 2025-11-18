// Code generated using sure/sure_pkg_gen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/sure
// Generated from: testing.go:1934 -> testing.tRunner
// ========== SURE:DO-NOT-EDIT-SECTION:END ==========

package jsonozh

import (
	"github.com/go-zwbc/jsonzh"
	"github.com/yyle88/sure"
)

func M编码(object any) []byte {
	res0, err := jsonzh.M编码(object)
	sure.Omit(err)
	return res0
}

func U解码[T any](data []byte) *T {
	res0, err := jsonzh.U解码[T](data)
	sure.Omit(err)
	return res0
}

func M编码s(object any) string {
	res0, err := jsonzh.M编码s(object)
	sure.Omit(err)
	return res0
}

func U解码s[T any](data string) *T {
	res0, err := jsonzh.U解码s[T](data)
	sure.Omit(err)
	return res0
}
