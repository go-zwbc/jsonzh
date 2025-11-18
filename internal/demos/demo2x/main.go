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
