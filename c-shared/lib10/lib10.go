package main

import (
	"fmt"
)

import "C"

//export lib10_f
func lib10_f() {
	fmt.Println("lib10.f")
}

func main() {}
