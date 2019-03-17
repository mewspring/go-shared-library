package main

import (
	"fmt"
)

import "C"

//export lib1_f
func lib1_f() {
	fmt.Println("lib1.f")
}

func main() {}
