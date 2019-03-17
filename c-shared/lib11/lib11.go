package main

import (
	"fmt"
)

import "C"

//export lib11_f
func lib11_f() {
	fmt.Println("lib11.f")
}

func main() {}
