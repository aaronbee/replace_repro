package main

import (
	"fmt"

	"golang.org/x/tools/go/analysis/passes/printf"
)

func main() {
	fmt.Println(printf.KindPrintf) // compile error if wrong dependency used
}
