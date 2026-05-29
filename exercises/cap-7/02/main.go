package main

import "fmt"

func main() {
	bornYear := 1990
	limitYear := 2026

	for bornYear <= limitYear {
		fmt.Println("The person was born in", bornYear)
		bornYear++
	}
}
