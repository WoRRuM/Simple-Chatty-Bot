package main

import "fmt"

func main() {
	// DO NOT delete or modify the code block below!
	var s1 = []int{12, 23, 34}
	var s2 = []int{45, 56, 67}

	s2 = append(s1, s2...)
	// Make a slice of strings with the length 'y' below:

	fmt.Println(s2) // print the 'len' of the slice here!
}
