package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func solve(s []string, x string, i int) []string {
    // Write the code to make a copy of the 's' slice below:
	c := []string
	c = append(
		s,
	)

	// Assign x to the element at the 'i' index of the 'c' slice:
	c[i] = x

	return c // Return the copy of the slice here!

	/*	var c []string
		c = append(s)

		// Assign x to the element at the 'i' index of the 'c' slice:
		c[i] = x

		return c */
}

// DO NOT delete or modify the contents of the main function!
func main() {
	var i int
	var x string

	s := make([]string, 1)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s[i], x = scanner.Text(), scanner.Text()

	n := solve(s, x, i)
	if !reflect.DeepEqual(s, n) {
		log.Fatal("You didn't assign 'x' to the element at the 'i' index of the slice!")
	}
	fmt.Println(n)
}