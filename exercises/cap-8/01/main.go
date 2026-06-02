package main

import "fmt"

func main() {

	x := []int{1, 2, 3}

	slice := x[:]

	fmt.Println(slice)

	for i, v := range x {
		fmt.Println("Indice: ", i, "valor: ", v)

	}
}
