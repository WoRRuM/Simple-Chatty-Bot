package main

import "fmt"

func main() {
	var p = new(int)
	var x int
	fmt.Scan(&x)

	// Write your code below:
	*p = x

	fmt.Println(*p)
}
