package main

import (
	"fmt"
)

import "C"

//export lib3_f
func lib3_f() {
	fmt.Println("lib3.f")
}

func main() {}
