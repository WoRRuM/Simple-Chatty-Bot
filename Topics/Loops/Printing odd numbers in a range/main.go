package main

import "fmt"

func main() {
	for i := 2; i <= 1023; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
