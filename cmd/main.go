package main

import (
	"fmt"
	chunk "pgnme/internal/chunk/chunk_type"
)

func main() {
	s := "RuST"
	fmt.Printf("%v\n", []byte(s))
	chunk.FromString("hello")
}
