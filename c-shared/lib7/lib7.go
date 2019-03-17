package main

import (
	"fmt"
)

import "C"

//export lib7_f
func lib7_f() {
	fmt.Println("lib7.f")
}

func main() {}
