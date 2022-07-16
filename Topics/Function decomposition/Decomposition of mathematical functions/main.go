package main

import "fmt"

// getHalf prints the result divided by 2
func getHalf(result int) {
	fmt.Println("The result is", result/2)
}

// halfSum sums a and b and then calls getHalf to print the result divided by 2
func halfSum(a int, b int) {
	getHalf(a + b)
}

// halfSubtract substracts b from a then calls getHalf to print the result divided by 2
func halfSubstract(a int, b int) {
	getHalf(a - b)
}

// DO NOT delete or modify the code within the main function!
func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	halfSum(a, b)
	halfSubstract(a, b)
}
