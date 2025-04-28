package jsonzh_test

import (
	"testing"

	"github.com/go-zwbc/jsonzh"
	"github.com/stretchr/testify/require"
)

type Person struct {
	N姓名 string `json:"name"`
	A年龄 int    `json:"age"`
	E邮箱 string `json:"eddress,omitempty"`
}

func TestM编码(t *testing.T) {
	// 测试用例1：正常结构体序列化
	t.Run("NormalStruct", func(t *testing.T) {
		input := Person{N姓名: "Alice", A年龄: 30, E邮箱: "alice@example.com"}
		expected := []byte(`{"name":"Alice","age":30,"eddress":"alice@example.com"}`)

		data, err := jsonzh.M编码(input)
		require.NoError(t, err, "M编码 should not return error")
		require.JSONEq(t, string(expected), string(data), "M编码 output should match expected JSON")
	})

	// 测试用例2：空结构体序列化
	t.Run("EmptyStruct", func(t *testing.T) {
		input := Person{}
		expected := []byte(`{"name":"","age":0}`)

		data, err := jsonzh.M编码(input)
		require.NoError(t, err, "M编码 should not return error for empty struct")
		require.JSONEq(t, string(expected), string(data), "M编码 output should match expected JSON")
	})

	// 测试用例3：不可序列化的输入（如 channel）
	t.Run("InvalidInput", func(t *testing.T) {
		input := make(chan int)
		_, err := jsonzh.M编码(input)
		require.Error(t, err, "M编码 should return error for invalid input")
	})
}

func TestU解码(t *testing.T) {
	// 测试用例1：正常 JSON 解析
	t.Run("NormalJSON", func(t *testing.T) {
		input := []byte(`{"name":"Bob","age":25,"eddress":"bob@example.com"}`)
		expected := Person{N姓名: "Bob", A年龄: 25, E邮箱: "bob@example.com"}

		result, err := jsonzh.U解码[Person](input)
		require.NoError(t, err, "U解码 should not return error")
		require.Equal(t, &expected, result, "U解码 output should match expected struct")
	})

	// 测试用例2：空 JSON 解析
	t.Run("EmptyJSON", func(t *testing.T) {
		input := []byte(`{}`)
		expected := Person{}

		result, err := jsonzh.U解码[Person](input)
		require.NoError(t, err, "U解码 should not return error for empty JSON")
		require.Equal(t, &expected, result, "U解码 output should match expected struct")
	})

	// 测试用例3：无效 JSON
	t.Run("InvalidJSON", func(t *testing.T) {
		input := []byte(`{invalid json}`)
		_, err := jsonzh.U解码[Person](input)
		require.Error(t, err, "U解码 should return error for invalid JSON")
	})
}

func TestM编码s(t *testing.T) {
	// 测试用例1：正常结构体序列化为字符串
	t.Run("NormalStruct", func(t *testing.T) {
		input := Person{N姓名: "Charlie", A年龄: 40}
		expected := `{"name":"Charlie","age":40}`

		data, err := jsonzh.M编码s(input)
		require.NoError(t, err, "M编码s should not return error")
		require.JSONEq(t, expected, data, "M编码s output should match expected JSON string")
	})

	// 测试用例2：空结构体序列化
	t.Run("EmptyStruct", func(t *testing.T) {
		input := Person{}
		expected := `{"name":"","age":0}`

		data, err := jsonzh.M编码s(input)
		require.NoError(t, err, "M编码s should not return error for empty struct")
		require.JSONEq(t, expected, data, "M编码s output should match expected JSON string")
	})

	// 测试用例3：不可序列化的输入
	t.Run("InvalidInput", func(t *testing.T) {
		input := make(chan int)
		_, err := jsonzh.M编码s(input)
		require.Error(t, err, "M编码s should return error for invalid input")
	})
}

func TestU解码s(t *testing.T) {
	// 测试用例1：正常 JSON 字符串解析
	t.Run("NormalJSON", func(t *testing.T) {
		input := `{"name":"Dave","age":35,"eddress":"dave@example.com"}`
		expected := Person{N姓名: "Dave", A年龄: 35, E邮箱: "dave@example.com"}

		result, err := jsonzh.U解码s[Person](input)
		require.NoError(t, err, "U解码s should not return error")
		require.Equal(t, &expected, result, "U解码s output should match expected struct")
	})

	// 测试用例2：空 JSON 字符串解析
	t.Run("EmptyJSON", func(t *testing.T) {
		input := `{}`
		expected := Person{}

		result, err := jsonzh.U解码s[Person](input)
		require.NoError(t, err, "U解码s should not return error for empty JSON")
		require.Equal(t, &expected, result, "U解码s output should match expected struct")
	})

	// 测试用例3：无效 JSON 字符串
	t.Run("InvalidJSON", func(t *testing.T) {
		input := `{invalid json}`
		_, err := jsonzh.U解码s[Person](input)
		require.Error(t, err, "U解码s should return error for invalid JSON")
	})

	// 测试用例4：空字符串
	t.Run("EmptyString", func(t *testing.T) {
		input := ""
		_, err := jsonzh.U解码s[Person](input)
		require.Error(t, err, "U解码s should return error for empty string")
	})
}
