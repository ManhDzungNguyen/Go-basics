package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("Go version: %s\n", runtime.Version())
}
