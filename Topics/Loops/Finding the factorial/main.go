package main

import (
	"fmt"
)

func main() {
	var n int

	var fact int = 1

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fact = fact * i
	}

	fmt.Printf("%d", fact)
}
