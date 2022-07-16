package main

import "fmt"

func whatName(name string) string {
	return name + " is learning how to call functions!"
}

func main() {
	// Do not change this two lines of code
	var userName string
	fmt.Scanf("%s", &userName)
	fmt.Println(whatName(userName))
	// call the function directly, or within a fmt.Println statement here.
}
