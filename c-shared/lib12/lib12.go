package main

import (
	"fmt"
)

import "C"

//export lib12_f
func lib12_f() {
	fmt.Println("lib12.f")
}

func main() {}
