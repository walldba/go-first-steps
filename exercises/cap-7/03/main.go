package main

import "fmt"

func main() {
	age := 20
	minAge := 18

	if age <= minAge {
		fmt.Println("You can not drive")
	} else {
		fmt.Println("You can drive")
	}
}
