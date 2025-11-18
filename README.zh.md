[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-zwbc/jsonzh/release.yml?branch=main&label=BUILD)](https://github.com/go-zwbc/jsonzh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-zwbc/jsonzh)](https://pkg.go.dev/github.com/go-zwbc/jsonzh)
[![Coverage Status](https://img.shields.io/coveralls/github/go-zwbc/jsonzh/main.svg)](https://coveralls.io/github/go-zwbc/jsonzh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25%2B-lightgrey.svg)](https://github.com/go-zwbc/jsonzh)
[![GitHub Release](https://img.shields.io/github/release/go-zwbc/jsonzh.svg)](https://github.com/go-zwbc/jsonzh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zwbc/jsonzh)](https://goreportcard.com/report/github.com/go-zwbc/jsonzh)

# jsonzh

使用中文函数名和增强错误处理的轻量级 Go JSON 编解码包

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 主要特性

🎯 **中文函数名**: 基于 Go 标准 JSON 库的直观中文命名封装
🔒 **增强错误处理**: 使用 `github.com/pkg/errors` 提供丰富的错误上下文和堆栈跟踪
🎨 **泛型类型支持**: 使用 Go 泛型 `[T any]` 实现完全类型安全的操作
📊 **双格式支持**: 同时支持字节切片和字符串格式的编解码
🌐 **多种错误策略**: 三个变体包提供不同的错误处理方式

## 安装

```bash
go get github.com/go-zwbc/jsonzh
```

## 快速开始

### 基础用法

使用 `M编码` 和 `U解码` 进行字节格式的 JSON 编解码。

```go
package main

import (
	"fmt"

	"github.com/go-zwbc/jsonzh"
	"github.com/yyle88/rese"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Status string `json:"status"`
}

func main() {
	// 创建一个 person 对象
	person := Person{
		Name:   "Alice",
		Age:    25,
		Status: "Active",
	}

	// 编码为 JSON 字节
	data := rese.V1(jsonzh.M编码(person))
	fmt.Printf("JSON 字节: %s\n", string(data))

	// 从 JSON 字节解码
	decoded := rese.P1(jsonzh.U解码[Person](data))
	fmt.Printf("解码结果: Name=%s, Age=%d, Status=%s\n", decoded.Name, decoded.Age, decoded.Status)
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 基于字符串的操作

使用 `M编码s` 和 `U解码s` 进行字符串格式的 JSON 编解码。

```go
package main

import (
	"fmt"

	"github.com/go-zwbc/jsonzh"
	"github.com/yyle88/rese"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	// 创建一个 product 对象
	product := Product{
		ID:    1001,
		Name:  "笔记本电脑",
		Price: 5999.99,
	}

	// 编码为 JSON 字符串
	jsonString := rese.C1(jsonzh.M编码s(product))
	fmt.Printf("JSON 字符串: %s\n", jsonString)

	// 从 JSON 字符串解码
	decoded := rese.P1(jsonzh.U解码s[Product](jsonString))
	fmt.Printf("解码结果: ID=%d, Name=%s, Price=%.2f\n", decoded.ID, decoded.Name, decoded.Price)
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

### 必须错误处理

使用 `jsonmzh` 处理不可接受错误、需要立即 panic 的操作。

```go
package main

import (
	"fmt"

	"github.com/go-zwbc/jsonzh/sure/jsonmzh"
	"github.com/yyle88/must"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Mode string `json:"mode"`
}

func main() {
	// 创建一个 config 对象
	config := Config{
		Host: "localhost",
		Port: 8080,
		Mode: "production",
	}

	// 编码为 JSON 字节（出错时 panic）
	data := must.Have(jsonmzh.M编码(config))
	fmt.Printf("JSON 字节: %s\n", string(data))

	// 从 JSON 字节解码（出错时 panic）
	decoded := must.Full(jsonmzh.U解码[Config](data))
	fmt.Printf("解码结果: Host=%s, Port=%d, Mode=%s\n", decoded.Host, decoded.Port, decoded.Mode)
}
```

⬆️ **源码:** [源码](internal/demos/demo3x/main.go)

### 使用错误处理变体

该包提供三个子包，提供不同的错误处理策略：

#### 1. jsonszh - 软错误处理（记录错误）

```go
import "github.com/go-zwbc/jsonzh/sure/jsonszh"

// 出错时返回默认值/零值并记录错误
data := jsonszh.M编码(person)     // 出错时返回空 []byte
result := jsonszh.U解码[Person](data) // 出错时返回 nil
```

#### 2. jsonmzh - 必须错误处理（出错时 panic）

```go
import "github.com/go-zwbc/jsonzh/sure/jsonmzh"

// 出错时 panic - 用于失败不可接受的场景
data := jsonmzh.M编码(person)     // 编码错误时 panic
result := jsonmzh.U解码[Person](data) // 解码错误时 panic
```

#### 3. jsonozh - 忽略错误处理（静默忽略错误）

```go
import "github.com/go-zwbc/jsonzh/sure/jsonozh"

// 静默忽略错误 - 用于失败可接受的场景
data := jsonozh.M编码(person)     // 出错时返回空 []byte
result := jsonozh.U解码[Person](data) // 出错时返回 nil
```

## API 参考

### 主包函数

#### 基于字节的操作

- `M编码(object any) ([]byte, error)` - 将任意对象编码为 JSON 字节
- `U解码[T any](data []byte) (*T, error)` - 将 JSON 字节解码为类型化对象

#### 基于字符串的操作

- `M编码s(object any) (string, error)` - 将任意对象编码为 JSON 字符串
- `U解码s[T any](data string) (*T, error)` - 将 JSON 字符串解码为类型化对象

### 错误处理变体

所有四个主要函数在三个变体包中都可用：

| 包名 | 错误策略 | 使用场景 |
|------|---------|---------|
| `jsonszh` | 软处理（记录） | 开发、调试 |
| `jsonmzh` | 必须（panic） | 关键操作 |
| `jsonozh` | 忽略（静默） | 非关键操作 |

## 错误上下文

该包的所有错误都包含增强的上下文信息：

```go
person, err := jsonzh.U解码[Person](invalidData)
if err != nil {
    // 错误包含上下文："解码错误: <原始错误>"
    // 加上 github.com/pkg/errors 的完整堆栈跟踪
    fmt.Printf("%+v\n", err) // 打印带堆栈跟踪的错误
}
```

## 函数名含义

- `M编码` - Marshal/Encode（编码）
- `U解码` - Unmarshal/Decode（解码）
- 后缀 `s` - 字符串变体（如 `M编码s`、`U解码s`）

## 安全性

该包使用 Go 标准库 `encoding/json`，提供安全的 JSON 解析和生成。编码时所有用户输入都会被正确转义，以防止 JSON 注入攻击。

## License

[LICENSE](LICENSE)
