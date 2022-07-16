package main

import "fmt"

func main() {
	var age int
	fmt.Scanf("%d", &age)

	if age <= 14 {
		fmt.Println("Toy Story 4")
	} else if age <= 18 && age >= 15 {
		fmt.Println("The Matrix")
	} else if age <= 25 && age >= 19 {
		fmt.Println("John Wick")
	} else if age <= 35 && age >= 26 {
		fmt.Println("Constantine")
	} else {
		fmt.Println("Speed")
	}
}
