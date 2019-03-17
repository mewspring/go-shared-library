package main

import (
	"fmt"
)

import "C"

//export lib6_f
func lib6_f() {
	fmt.Println("lib6.f")
}

func main() {}
