package main

import (
	"fmt"
)

// demo code
func Print[T any](value T) {
	fmt.Println(value)
}

func main() {
	Print("Hello, Generics!")
	Print(42)
	Print(3.14)
}