package main

import (
	"fmt"
	"log"
	"reflect"
)

func solve(s1, s2 []string) []string {
	s1 = append(s1, s2...)

	return s1 // return the appended slice here!
}

// DO NOT delete or modify the contents of the main function!
func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	s1, s2 := make([]string, n), make([]string, m)

	for i := range s1 {
		fmt.Scanln(&s1[i])
	}

	for i := range s2 {
		fmt.Scanln(&s2[i])
	}

	if !reflect.DeepEqual(append(s1, s2...), solve(s1, s2)) {
		log.Fatal("slices were not properly appended")
	}
	fmt.Println(solve(s1, s2))
}
