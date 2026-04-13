package main

import (
	"fmt"

	"github.com/oapi-codegen/nullable"
)

func main() {
	p := struct {
		N nullable.Nullable[int]
	}{}

	p.N = nullable.NewNullNullable[int]()

	fmt.Printf("Specified: %v\n", p.N.IsSpecified())
	fmt.Printf("Null: %v\n", p.N.IsNull())
}
