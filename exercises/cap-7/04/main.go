package main

import "fmt"

func main() {
	age := 20
	minAge := 18

	switch {
	case age >= minAge:
		fmt.Println("You can drive")
	default:
		fmt.Println("You can not drive")
	}
}
