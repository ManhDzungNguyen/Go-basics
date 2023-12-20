package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]map[string]int)
	a["l1.a"] = make(map[string]int)
	a["l1.a"]["l2.a"] = 42
	a["l1.a"]["l2.b"] = 68

	fmt.Println(a)
	fmt.Println(a["l1.a"])
	fmt.Println(a["l1.a"]["l2.a"])
	fmt.Println(a["l1.b"])
	fmt.Println(a["l1.b"]["l2.b"])
}
