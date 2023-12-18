package main

import (
	"Basics/calculate"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	sum := calculate.Total(200, 40)
	fmt.Println(sum)

	//------

	var a = make(map[string]map[string]int)
	a["abc"] = make(map[string]int)
	a["abc"]["def"] = 2
	fmt.Println(a["de"]["abc"])

}
