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
