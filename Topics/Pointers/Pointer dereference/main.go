package main

import "fmt"

func solve(p *int) int {
	var v = p // assign the value of the pointer 'p' to the variable 'v'

	return *v // return the value of the variable 'v'
}

// DO NOT change the contents of the main function!
func main() {
	var i = new(int)
	fmt.Scan(i)

	var v = solve(i)
	fmt.Println(v)
}
