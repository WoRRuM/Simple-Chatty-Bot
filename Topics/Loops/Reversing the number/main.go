package main

import "fmt"

func reverseInt(n int) int {
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%d", reverseInt(n))
}
