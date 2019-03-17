package main

import (
	"fmt"
)

import "C"

//export lib2_f
func lib2_f() {
	fmt.Println("lib2.f")
}

func main() {}
