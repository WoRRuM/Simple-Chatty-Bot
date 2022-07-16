package main

import "fmt"

func main() {
	/*	var number int
		fmt.Scanf("%d", &number)

		if number > 0 {
			fmt.Println("Positive!")
		} else if number < 0 {
			fmt.Println("Negative!")
		} else {
			fmt.Println("Zero!")
		}*/
	var number int
	number, err := fmt.Scanf("%d", &number)
	if err != nil {
		return
	}

	if number > 0 {
		fmt.Println("Positive!")
	} else if number < 0 {
		fmt.Println("Negative!")
	} else {
		fmt.Println("Zero!")
	}
}
