package main

import (
	"fmt"
)

import "C"

//export lib5_f
func lib5_f() {
	fmt.Println("lib5.f")
}

func main() {}
