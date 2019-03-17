package main

import (
	"fmt"
)

import "C"

//export lib4_f
func lib4_f() {
	fmt.Println("lib4.f")
}

func main() {}
