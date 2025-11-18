[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-zwbc/jsonzh/release.yml?branch=main&label=BUILD)](https://github.com/go-zwbc/jsonzh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-zwbc/jsonzh)](https://pkg.go.dev/github.com/go-zwbc/jsonzh)
[![Coverage Status](https://img.shields.io/coveralls/github/go-zwbc/jsonzh/main.svg)](https://coveralls.io/github/go-zwbc/jsonzh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25%2B-lightgrey.svg)](https://github.com/go-zwbc/jsonzh)
[![GitHub Release](https://img.shields.io/github/release/go-zwbc/jsonzh.svg)](https://github.com/go-zwbc/jsonzh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zwbc/jsonzh)](https://goreportcard.com/report/github.com/go-zwbc/jsonzh)

# jsonzh

Lightweight Go package with JSON encoding/decoding, using Chinese function names and enhanced errors handling

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

🎯 **Chinese Function Names**: Intuitive Chinese-named wrappers around Go's standard JSON package
🔒 **Enhanced Errors Handling**: Rich errors context with stack traces using `github.com/pkg/errors`
🎨 **Generic Type Support**: Type-safe operations with Go generics `[T any]`
📊 **Format Support**: Both byte slice and string-based encoding/decoding
🌐 **Multiple Errors Strategies**: Three variant packages with different errors handling approaches

## Installation

```bash
go get github.com/go-zwbc/jsonzh
```

## Quick Start

### Basic Usage

Encode and decode JSON using byte format with `M编码` and `U解码`.

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
	// Create a person object
	person := Person{
		Name:   "Alice",
		Age:    25,
		Status: "Active",
	}

	// Encode to JSON bytes
	data := rese.V1(jsonzh.M编码(person))
	fmt.Printf("JSON bytes: %s\n", string(data))

	// Decode from JSON bytes
	decoded := rese.P1(jsonzh.U解码[Person](data))
	fmt.Printf("Decoded: Name=%s, Age=%d, Status=%s\n", decoded.Name, decoded.Age, decoded.Status)
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### String-Based Operations

Encode and decode JSON using string format with `M编码s` and `U解码s`.

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
	// Create a product object
	product := Product{
		ID:    1001,
		Name:  "Laptop",
		Price: 5999.99,
	}

	// Encode to JSON string
	jsonString := rese.C1(jsonzh.M编码s(product))
	fmt.Printf("JSON string: %s\n", jsonString)

	// Decode from JSON string
	decoded := rese.P1(jsonzh.U解码s[Product](jsonString))
	fmt.Printf("Decoded: ID=%d, Name=%s, Price=%.2f\n", decoded.ID, decoded.Name, decoded.Price)
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

### Must Errors Handling

Use `jsonmzh` when operations with errors being unacceptable and should cause immediate panic.

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
	// Create a config object
	config := Config{
		Host: "localhost",
		Port: 8080,
		Mode: "production",
	}

	// Encode to JSON bytes (panics on failure)
	data := must.Have(jsonmzh.M编码(config))
	fmt.Printf("JSON bytes: %s\n", string(data))

	// Decode from JSON bytes (panics on failure)
	decoded := must.Full(jsonmzh.U解码[Config](data))
	fmt.Printf("Decoded: Host=%s, Port=%d, Mode=%s\n", decoded.Host, decoded.Port, decoded.Mode)
}
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

### Using Errors Handling Variants

The package provides three subpackages with different errors handling strategies:

#### 1. jsonszh - Soft Errors Handling (Logs errors)

```go
import "github.com/go-zwbc/jsonzh/sure/jsonszh"

// Returns default/zero values on failures and logs them
data := jsonszh.M编码(person)     // Returns []byte with zero-length on failure
result := jsonszh.U解码[Person](data) // Returns nil on failure
```

#### 2. jsonmzh - Must Errors Handling (Panics on failures)

```go
import "github.com/go-zwbc/jsonzh/sure/jsonmzh"

// Panics on failures - use when being unacceptable
data := jsonmzh.M编码(person)     // Panics on encoding failure
result := jsonmzh.U解码[Person](data) // Panics on decoding failure
```

#### 3. jsonozh - Omit Errors Handling (Ignores errors in silence)

```go
import "github.com/go-zwbc/jsonzh/sure/jsonozh"

// Ignores errors in silence - use when being acceptable
data := jsonozh.M编码(person)     // Returns []byte with zero-length on failure
result := jsonozh.U解码[Person](data) // Returns nil on failure
```

## API Reference

### Main Package Functions

#### Byte-Based Operations

- `M编码(object any) ([]byte, error)` - Encode any object to JSON bytes
- `U解码[T any](data []byte) (*T, error)` - Decode JSON bytes to typed object

#### String-Based Operations

- `M编码s(object any) (string, error)` - Encode any object to JSON string
- `U解码s[T any](data string) (*T, error)` - Decode JSON string to typed object

### Errors Handling Variants

The following table shows that the main functions are available in three variant packages:

| Package | Errors Handling | Use Case |
|---------|---------------|----------|
| `jsonszh` | Soft (logs) | Development, debugging |
| `jsonmzh` | Must (panics) | Operations being critical |
| `jsonozh` | Omit (silent) | Operations being non-critical |

## Errors Context

The errors from this package includes enhanced context:

```go
person, err := jsonzh.U解码[Person](invalidData)
if err != nil {
    // Errors includes context: "解码错误: <the first errors>"
    // With stack trace from github.com/pkg/errors
    fmt.Printf("%+v\n", err) // Prints errors with stack trace
}
```

## Function Name Meanings

- `M编码` - Marshal/Encode (编码 = encoding)
- `U解码` - Unmarshal/Decode (解码 = decoding)
- Suffix `s` - String variant (e.g., `M编码s`, `U解码s`)

## Safe Design

This package uses Go's standard `encoding/json` package, which provides safe JSON parsing and generation. When encoding to prevent JSON injection attacks, this package escapes the input from the users.

## License

[LICENSE](LICENSE)
