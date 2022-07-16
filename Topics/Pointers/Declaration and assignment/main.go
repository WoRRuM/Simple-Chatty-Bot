package main

import "fmt"

func main() {
	var p = new(string)
	*p = "my string"

	fmt.Println(*p) // print your pointer variable here!
}
